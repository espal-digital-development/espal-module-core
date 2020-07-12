package forgot

import (
	"net/http"
	"strings"

	"github.com/espal-digital-development/espal-core/config"
	"github.com/espal-digital-development/espal-core/mailer"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/user"
	"github.com/espal-digital-development/espal-core/text"
	"github.com/espal-digital-development/espal-core/validators/forms/account/password/forgot"
	page "github.com/espal-digital-development/espal-module-core/pages/account/password/forgot"
	"github.com/juju/errors"
)

const hashLength = 72

// Route processor.
type Route struct {
	configService               config.Config
	mailerService               mailer.Engine
	userStore                   user.Store
	forgotPasswordFormValidator forgot.Factory
	forgotPageFactory           page.Factory
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	if context.IsLoggedIn() {
		context.Redirect("/", http.StatusTemporaryRedirect)
		return
	}

	language, err := context.GetLanguage()
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	form, err := r.forgotPasswordFormValidator.New(language)
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	defer form.Close()
	isSubmitted, isValid, err := form.Submit(context)
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	if isSubmitted && isValid {
		user, ok, err := r.userStore.GetOneActiveByEmail(form.FormFieldValue("email"))
		if err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
		if !ok {
			context.Redirect("/ForgotPasswordSucceeded", http.StatusTemporaryRedirect)
			return
		}

		hash := text.RandomString(hashLength)
		if err := r.userStore.SetPasswordResetHashForUser(user.ID(), hash); err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}

		var name string
		if user.FirstName() != nil {
			name = *user.FirstName()
		} else {
			name = user.Email()
		}
		// TODO :: 77777 Mail Templates from file and/or database with qtpl and ftpl
		body := strings.Join([]string{
			"<p>Hello ", name,
			",</p><p>Password Forget Explanation</p>",
			"<p><a href=\"",
			strings.Join([]string{context.GetDomain().HostWithProtocol(), "/PasswordRecovery?h=", hash}, ""),
			"\">Recover here</a></p>",
			"<p>Regards,</p><p>- Espal</p>",
		}, "")

		message := r.mailerService.NewMessage()
		message.SetHeader("From", r.configService.EmailNoReplyAddress())
		message.SetHeader("To", user.Email())
		message.SetHeader("Subject", context.Translate("passwordResetRequest"))
		message.SetBody("text/html", body)

		// Don't wait for the mail
		go func(forgotRoute *Route, routeContext contexts.Context, msg mailer.Data) {
			if err := forgotRoute.mailerService.Send(msg); err != nil {
				routeContext.RenderInternalServerError(errors.Trace(err))
				return
			}
		}(r, context, message)

		context.Redirect("/ForgotPasswordSucceeded", http.StatusTemporaryRedirect)
		return
	}

	r.forgotPageFactory.NewPage(context, form.View()).Render()
}

// New returns a new instance of Route.
func New(configService config.Config, mailerService mailer.Engine, userStore user.Store,
	forgotPasswordFormValidator forgot.Factory, forgotPageFactory page.Factory) *Route {
	return &Route{
		configService:               configService,
		mailerService:               mailerService,
		userStore:                   userStore,
		forgotPasswordFormValidator: forgotPasswordFormValidator,
		forgotPageFactory:           forgotPageFactory,
	}
}
