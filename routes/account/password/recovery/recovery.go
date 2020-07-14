package recovery

import (
	"net/http"

	"github.com/espal-digital-development/espal-core/config"
	"github.com/espal-digital-development/espal-core/repositories/regularexpressions"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/user"
	"github.com/espal-digital-development/espal-module-core/forms/account/password/recovery"
	page "github.com/espal-digital-development/espal-module-core/pages/account/password/recovery"
	"github.com/juju/errors"
	"golang.org/x/crypto/bcrypt"
)

// Route processor.
type Route struct {
	configService                config.Config
	regularExpressionsRepository regularexpressions.Repository
	userStore                    user.Store
	recoveryFormValidator        recovery.Factory
	recoveryPageFactory          page.Factory
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	if context.IsLoggedIn() {
		context.Redirect("/", http.StatusTemporaryRedirect)
		return
	}

	hash := context.QueryValue("h")
	if hash == "" || !r.regularExpressionsRepository.GetPasswordRecoveryhash().MatchString(hash) {
		context.RenderBadRequest()
		return
	}

	language, err := context.GetLanguage()
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}

	form, err := r.recoveryFormValidator.New(language)
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
		newPassword, err := bcrypt.GenerateFromPassword([]byte(form.FormFieldValue("newPassword")),
			r.configService.SecurityBcryptRounds())
		if err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}

		var resetCount uint8
		if c := form.GetPasswordResetCount(); c != nil {
			resetCount = *c
		}
		resetCount++

		err = r.userStore.RecoverWithNewPassword(form.GetUserID(), newPassword, resetCount)
		if err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}

		// TODO :: 7 Test if this will actually be instantly logged in after redirect, or only after the page load afterwards
		err = context.Login(form.GetUserID(), false)
		if err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
		context.Redirect("/", http.StatusTemporaryRedirect)
		return
	}

	r.recoveryPageFactory.NewPage(context, form.View()).Render()
}

// New returns a new instance of Route.
func New(configService config.Config, regularExpressionsRepository regularexpressions.Repository, userStore user.Store,
	recoveryFormValidator recovery.Factory, recoveryPageFactory page.Factory) *Route {
	return &Route{
		configService:                configService,
		regularExpressionsRepository: regularExpressionsRepository,
		userStore:                    userStore,
		recoveryFormValidator:        recoveryFormValidator,
		recoveryPageFactory:          recoveryPageFactory,
	}
}
