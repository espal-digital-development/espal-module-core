package createupdate

import (
	"fmt"
	"net/http"

	"github.com/espal-digital-development/espal-core/database/entitymutators"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/user"
	"github.com/espal-digital-development/espal-core/stores/user/contact"
	"github.com/espal-digital-development/espal-module-core/forms/admin/user/contact/createupdate"
	page "github.com/espal-digital-development/espal-module-core/pages/admin/user/contact/createupdate"
	"github.com/juju/errors"
)

// Route processor.
type Route struct {
	entityMutatorsFactory     entitymutators.Factory
	userStore                 user.Store
	userContactStore          contact.Store
	createUpdateFormValidator createupdate.Factory
	createUpdatePageFactory   page.Factory
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	user, ok, err := r.userStore.GetOne(context.QueryValue("userid"))
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	if !ok {
		context.RenderBadRequest()
		return
	}

	id := context.QueryValue("id")

	if (id == "" && !context.HasUserRightOrForbid("CreateUserContact")) ||
		(id != "" && !context.HasUserRightOrForbid("UpdateUserContact")) {
		return
	}

	userContact := &contact.Contact{}
	if id != "" {
		userContact, ok, err = r.userContactStore.GetOneByIDWithCreator(id)
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

	form, err := r.createUpdateFormValidator.New(userContact, language)
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
		userID := form.FieldValue("userID")
		entityMutator := r.entityMutatorsFactory.NewMutation(userContact, form, "User/Conttact")
		entityMutator.SetNullableString("userID", &userID, userContact.UserID())
		entityMutator.SetString("contactID", form.FieldPointerValue("contact"), userContact.ContactID())
		entityMutator.SetUint("sorting", form.FieldValueAsUint("sorting"), userContact.Sorting())
		entityMutator.SetNullableString("comments", form.FieldPointerValue("comment"), userContact.Comments())

		entityMutator.SetExtraURLQueryParams(fmt.Sprintf("userid=%s", user.ID()))
		entityMutator.SetCustomReturnPath(fmt.Sprintf("UserContact?id=%s", user.ID()))
		if err := entityMutator.Execute(context); err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}

		if redirectTo := entityMutator.RedirectURL(); redirectTo != "" {
			context.Redirect(context.AdminURL()+redirectTo, http.StatusTemporaryRedirect)
			return
		}

		userContact, ok, err = r.userContactStore.GetOneByIDWithCreator(entityMutator.GetInsertedOrUpdatedID())
		if err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
		if !ok {
			context.RenderNotFound()
			return
		}
	}

	// TODO :: 77 Pressing the action `cancel` acts strange and goes to a
	// broken URL (e.g. `/_PY0l/User/Contact/User/View?id=385509484458901505`).

	r.createUpdatePageFactory.NewPage(context, userContact, user, language, form.View(),
		context.GetAdminCreateUpdateTitle(id, "userContact")).Render()
}

// New returns a new instance of Route.
func New(entityMutatorsFactory entitymutators.Factory, userStore user.Store, userContactStore contact.Store,
	createUpdateFormValidator createupdate.Factory, createUpdatePageFactory page.Factory) *Route {
	return &Route{
		entityMutatorsFactory:     entityMutatorsFactory,
		userStore:                 userStore,
		userContactStore:          userContactStore,
		createUpdateFormValidator: createUpdateFormValidator,
		createUpdatePageFactory:   createUpdatePageFactory,
	}
}
