package createupdate

import (
	"net/http"
	"strings"

	"github.com/espal-digital-development/espal-core/config"
	"github.com/espal-digital-development/espal-core/database/entitymutators"
	"github.com/espal-digital-development/espal-core/routing/assethandler"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/storage"
	"github.com/espal-digital-development/espal-core/stores/user"
	"github.com/espal-digital-development/espal-core/stores/user/address"
	"github.com/espal-digital-development/espal-module-core/forms/admin/user/createupdate"
	page "github.com/espal-digital-development/espal-module-core/pages/admin/user/createupdate"
	"github.com/juju/errors"
	"golang.org/x/crypto/bcrypt"
)

// Route processor.
type Route struct {
	configService             config.Config
	assetHandler              assethandler.Handler
	entityMutatorsFactory     entitymutators.Factory
	assetsPublicFilesStorage  storage.Storage
	userStore                 user.Store
	userAddressStore          address.Store
	createUpdateFormValidator createupdate.Factory
	createUpdatePageFactory   page.Factory
}

// Handle route handler.
// nolint:funlen
func (r *Route) Handle(context contexts.Context) {
	id := context.QueryValue("id")

	if (id == "" && !context.HasUserRightOrForbid("CreateUser")) ||
		(id != "" && !context.HasUserRightOrForbid("UpdateUser")) {
		return
	}

	var err error
	var ok bool
	user := &user.User{}
	if id != "" {
		user, ok, err = r.userStore.GetOneByIDWithCreator(id)
		if err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
		if !ok {
			context.RenderNotFound()
			return
		}
	}

	language, err := context.GetLanguage()
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}

	form, err := r.createUpdateFormValidator.New(user, language)
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
		countryID := form.FieldValueAsUint16("country")
		currencies := strings.Join(form.FieldValues("currencies"), ",")

		entityMutator := r.entityMutatorsFactory.NewMutation(user, form, "User")
		entityMutator.SetBool("active", form.FieldValueAsBool("active"), user.Active())
		entityMutator.SetUint16("language", form.FieldValueAsUint16("language"), user.Language())
		entityMutator.SetNullableUint16("country", &countryID, user.Country())
		entityMutator.SetNullableString("firstName", form.FieldPointerValue("firstName"), user.FirstName())
		entityMutator.SetNullableString("surname", form.FieldPointerValue("surname"), user.Surname())
		if err := entityMutator.SetNullableTime("dateOfBirth", form.FieldPointerValue("dateOfBirth"),
			user.DateOfBirth()); err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
		entityMutator.SetString("email", form.FieldPointerValue("email"), user.Email())
		entityMutator.SetUint("priority", form.FieldValueAsUint("priority"), user.Priority())
		entityMutator.SetNullableString("defaultDeliveryAddressID", form.FieldPointerValue("defaultDeliveryAddress"),
			user.DefaultDeliveryAddressID())
		entityMutator.SetNullableString("defaultInvoiceAddressID", form.FieldPointerValue("defaultInvoiceAddress"),
			user.DefaultInvoiceAddressID())
		entityMutator.SetString("currencies", &currencies, user.Currencies())
		entityMutator.SetNullableString("biography", form.FieldPointerValue("biography"), user.Comments())
		entityMutator.SetNullableString("comments", form.FieldPointerValue("comments"), user.Comments())

		if form.FieldValue("password") != "" {
			encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(form.FieldValue("password")),
				r.configService.SecurityBcryptRounds())
			if err != nil {
				context.RenderInternalServerError(errors.Trace(err))
				return
			}
			encryptedPasswordString := string(encryptedPassword)
			entityMutator.SetString("password", &encryptedPasswordString, user.Password())
		}

		if form.HasAvatarUploaded() {
			if err := form.SaveAvatar(); err != nil {
				context.RenderInternalServerError(errors.Trace(err))
				return
			}

			fileName := form.AvatarFileName()
			entityMutator.SetNullableString("avatar", &fileName, user.Avatar())
			data, ok, err := r.assetsPublicFilesStorage.Get(form.AvatarSavedPath())
			if err != nil {
				context.RenderInternalServerError(errors.Trace(err))
				return
			}
			if !ok {
				context.RenderInternalServerError(
					errors.Errorf("saved avatar file wasn't found after seemily been correctly saved by the form"))
				return
			}
			if err := r.assetHandler.RegisterPublicFileRoute(form.AvatarSavedPath(), data); err != nil {
				context.RenderInternalServerError(errors.Trace(err))
				return
			}
		}

		if err := entityMutator.Execute(context); err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}

		user, ok, err = r.userStore.GetOneByIDWithCreator(entityMutator.GetInsertedOrUpdatedID())
		if err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
		if !ok {
			context.RenderNotFound()
			return
		}

		if redirectTo := entityMutator.RedirectURL(); redirectTo != "" {
			context.Redirect(context.AdminURL()+redirectTo, http.StatusTemporaryRedirect)
			return
		}
	}

	r.createUpdatePageFactory.NewPage(context, user, language, form.View(),
		context.GetAdminCreateUpdateTitle(id, "user")).Render()
}

// New returns a new instance of Route.
func New(configService config.Config, assetHandler assethandler.Handler,
	entityMutatorsFactory entitymutators.Factory, assetsPublicFilesStorage storage.Storage, userStore user.Store,
	userAddressStore address.Store, createUpdateFormValidator createupdate.Factory,
	createUpdatePageFactory page.Factory) *Route {
	return &Route{
		configService:             configService,
		assetHandler:              assetHandler,
		entityMutatorsFactory:     entityMutatorsFactory,
		assetsPublicFilesStorage:  assetsPublicFilesStorage,
		userStore:                 userStore,
		userAddressStore:          userAddressStore,
		createUpdateFormValidator: createUpdateFormValidator,
		createUpdatePageFactory:   createUpdatePageFactory,
	}
}
