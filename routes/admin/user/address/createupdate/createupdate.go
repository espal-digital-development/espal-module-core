package createupdate

import (
	"fmt"
	"net/http"

	"github.com/espal-digital-development/espal-core/database/entitymutators"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/user"
	"github.com/espal-digital-development/espal-core/stores/user/address"
	"github.com/espal-digital-development/espal-module-core/forms/admin/user/address/createupdate"
	page "github.com/espal-digital-development/espal-module-core/pages/admin/user/address/createupdate"
	"github.com/juju/errors"
)

// Route processor.
type Route struct {
	entityMutatorsFactory     entitymutators.Factory
	userStore                 user.Store
	userAddressStore          address.Store
	createUpdateFormValidator createupdate.Factory
	createUpdatePageFactory   page.Factory
}

// Handle route handler.
// nolint:funlen
func (r *Route) Handle(context contexts.Context) {
	// TODO :: 77777 This doesn't fetch (id issue?)
	user, ok, err := r.userStore.GetOne(context.QueryValue("id"))
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	if !ok {
		context.RenderBadRequest()
		return
	}

	id := context.QueryValue("id")

	if (id == "" && !context.HasUserRightOrForbid("CreateUserAddress")) ||
		(id != "" && !context.HasUserRightOrForbid("UpdateUserAddress")) {
		return
	}

	userAddress := &address.Address{}
	if id != "" {
		userAddress, ok, err = r.userAddressStore.GetOneByIDWithCreator(id)
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

	form, err := r.createUpdateFormValidator.New(userAddress, language)
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
		userID := user.ID()

		entityMutator := r.entityMutatorsFactory.NewMutation(userAddress, form, "User/Address")
		entityMutator.SetString("userID", &userID, userAddress.UserID())
		entityMutator.SetBool("active", form.FieldValueAsBool("active"), userAddress.Active())
		entityMutator.SetNullableString("firstName", form.FieldPointerValue("firstName"), userAddress.FirstName())
		entityMutator.SetNullableString("surname", form.FieldPointerValue("surname"), userAddress.Surname())
		entityMutator.SetString("street", form.FieldPointerValue("street"), userAddress.Street())
		entityMutator.SetString("number", form.FieldPointerValue("number"), userAddress.Number())
		entityMutator.SetNullableString("numberAddition", form.FieldPointerValue("numberAddition"),
			userAddress.NumberAddition())
		entityMutator.SetString("zipCode", form.FieldPointerValue("zipCode"), userAddress.ZipCode())
		entityMutator.SetString("city", form.FieldPointerValue("city"), userAddress.City())
		entityMutator.SetNullableUint16("country", &countryID, userAddress.Country())
		entityMutator.SetNullableString("phoneNumber", form.FieldPointerValue("phoneNumber"), userAddress.PhoneNumber())
		entityMutator.SetNullableString("email", form.FieldPointerValue("email"), userAddress.Email())

		entityMutator.SetExtraURLQueryParams(fmt.Sprintf("userid=%s", user.ID()))
		entityMutator.SetCustomReturnPath(fmt.Sprintf("User/View?id=%s", user.ID()))
		if err := entityMutator.Execute(context); err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}

		if redirectTo := entityMutator.RedirectURL(); redirectTo != "" {
			context.Redirect(context.AdminURL()+redirectTo, http.StatusTemporaryRedirect)
			return
		}

		userAddress, ok, err = r.userAddressStore.GetOneByIDWithCreator(entityMutator.GetInsertedOrUpdatedID())
		if err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
		if !ok {
			context.RenderNotFound()
			return
		}
	}

	r.createUpdatePageFactory.NewPage(context, userAddress, user, language, form.View(),
		context.GetAdminCreateUpdateTitle(id, "userAddress")).Render()
}

// New returns a new instance of Route.
func New(entityMutatorsFactory entitymutators.Factory, userStore user.Store, userAddressStore address.Store,
	createUpdateFormValidator createupdate.Factory, createUpdatePageFactory page.Factory) *Route {
	return &Route{
		entityMutatorsFactory:     entityMutatorsFactory,
		userStore:                 userStore,
		userAddressStore:          userAddressStore,
		createUpdateFormValidator: createUpdateFormValidator,
		createUpdatePageFactory:   createUpdatePageFactory,
	}
}
