package change

import (
	"net/http"

	"github.com/espal-digital-development/espal-core/config"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/user"
	"github.com/espal-digital-development/espal-module-core/forms/account/password/change"
	page "github.com/espal-digital-development/espal-module-core/pages/account/password/change"
	"github.com/juju/errors"
	"golang.org/x/crypto/bcrypt"
)

// Route processor.
type Route struct {
	configService               config.Config
	userStore                   user.Store
	changePasswordFormValidator change.Factory
	changePageFactory           page.Factory
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	if !context.IsLoggedIn() {
		context.Redirect("/Login", http.StatusTemporaryRedirect)
		return
	}

	user, ok, err := context.GetUser()
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	if !ok {
		context.RenderNotFound()
		return
	}

	language, err := context.GetLanguage()
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	form, err := r.changePasswordFormValidator.New(language, user)
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
		if err := r.userStore.SetPasswordForUser(user.ID(), newPassword); err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
		user.SetPassword(string(newPassword))

		if err := context.SetFlashSuccessMessage(context.Translate("yourPasswordHasBeenUpdated")); err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
	}

	r.changePageFactory.NewPage(context, form.View()).Render()
}

// New returns a new instance of Route.
func New(configService config.Config, userStore user.Store, changePasswordFormValidator change.Factory,
	changePageFactory page.Factory) *Route {
	return &Route{
		configService:               configService,
		userStore:                   userStore,
		changePasswordFormValidator: changePasswordFormValidator,
		changePageFactory:           changePageFactory,
	}
}
