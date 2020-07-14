package register

import (
	"net/http"
	"strings"

	"github.com/espal-digital-development/espal-core/config"
	"github.com/espal-digital-development/espal-core/mailer"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/user"
	"github.com/espal-digital-development/espal-module-core/forms/account/register"
	page "github.com/espal-digital-development/espal-module-core/pages/account/register"
	"github.com/juju/errors"
	"golang.org/x/crypto/bcrypt"
)

// Route processor.
type Route struct {
	configService         config.Config
	mailerService         mailer.Engine
	userStore             user.Store
	registerFormValidator register.Factory
	registerPageFactory   page.Factory
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	_, ok, err := context.GetUser()
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	if ok {
		context.Redirect("/", http.StatusTemporaryRedirect)
		return
	}

	language, err := context.GetLanguage()
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}

	form, err := r.registerFormValidator.New(language)
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
		// if dateOfBirth := form.Field("dateOfBirth").Value; dateOfBirth != "" {
		// 	// TODO :: Create a time.Time based on the Timestamp
		// 	// (YYYY-MM-DD or any variation of it)
		// 	// TODO :: Move inside the validator to a field.RenderedTime
		// 	// http://stackoverflow.com/questions/25845172/parsing-date-string-in-golang
		// 	// formattedDateOfBirth := nil
		// 	// user.DateOfBirth = &formattedDateOfBirth
		// }

		var firstName *string
		if firstNameValue := form.FormFieldValue("firstName"); firstNameValue != "" {
			firstName = &firstNameValue
		}

		var surname *string
		if surnameValue := form.FormFieldValue("surname"); surnameValue != "" {
			surname = &surnameValue
		}

		email := form.FormFieldValue("email")
		password := form.FormFieldValue("password")

		var name string
		if firstName != nil {
			name = *firstName
		} else {
			name = email
		}

		newPassword, err := bcrypt.GenerateFromPassword([]byte(password), r.configService.SecurityBcryptRounds())
		if err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}

		activationHash, err := r.userStore.Register(email, newPassword, firstName, surname, language.ID())
		if err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}

		// TODO :: On activation; add user to the correct groups too. The groups should
		// be determined by some setting based on CMS/Forum/Blog/B2C/B2B/ERP logic.

		// TODO :: 7 Use database-store based translated MailTemplate based on FastTemplate
		body := strings.Join([]string{
			"<p>Hello ", name,
			",</p><p>Account Activation Explanation</p>",
			"<p><a href=\"",
			strings.Join([]string{context.GetDomain().HostWithProtocol(), "/ActivateAccount?h=", activationHash}, ""),
			"\">Activate here</a></p>",
			"<p>Regards,</p><p>- Espal</p>",
		}, "")

		message := r.mailerService.NewMessage()
		message.SetHeader("From", r.configService.EmailNoReplyAddress())
		message.SetHeader("To", email)
		message.SetHeader("Subject", context.Translate("accountActivation"))
		message.SetBody("text/html", body)

		// Don't wait for the mail
		go func(registerRoute *Route, routeContext contexts.Context, msg mailer.Data) {
			if err := registerRoute.mailerService.Send(msg); err != nil {
				routeContext.RenderInternalServerError(errors.Trace(err))
				return
			}
		}(r, context, message)

		context.Redirect("/RegisterAccountSucceeded", http.StatusTemporaryRedirect)
		return
	}

	r.registerPageFactory.NewPage(context, form.View()).Render()
}

// New returns a new instance of Route.
func New(configService config.Config, mailerService mailer.Engine, userStore user.Store,
	registerFormValidator register.Factory, registerPageFactory page.Factory) *Route {
	return &Route{
		configService:         configService,
		mailerService:         mailerService,
		userStore:             userStore,
		registerFormValidator: registerFormValidator,
		registerPageFactory:   registerPageFactory,
	}
}
