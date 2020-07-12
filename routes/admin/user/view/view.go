package view

import (
	"fmt"

	"github.com/espal-digital-development/espal-core/pageactions"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/user"
	"github.com/espal-digital-development/espal-core/stores/user/address"
	"github.com/espal-digital-development/espal-core/stores/user/contact"
	page "github.com/espal-digital-development/espal-module-core/pages/admin/user/view"
	"github.com/juju/errors"
)

// Route processor.
type Route struct {
	userStore        user.Store
	userAddressStore address.Store
	userContactStore contact.Store
	viewPageFactory  page.Factory
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	if !context.HasUserRightOrForbid("ReadUser") {
		return
	}

	id := context.QueryValue("id")
	if id == "" {
		context.RenderNotFound()
		return
	}

	user, ok, err := r.userStore.GetOne(id)
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	if !ok {
		context.RenderNotFound()
		return
	}

	addresses, addressesOk, err := r.userAddressStore.ForUser(user.ID())
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	contacts, contactsOk, err := r.userContactStore.ForUser(user.ID())
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}

	var canUpdateAddress bool
	var canDeleteAddress bool
	var canUpdateContact bool
	var canDeleteContact bool
	addressesActions := pageactions.New(context, "UserAddress", addressesOk)
	addressesActions.AddCreateWithPath(fmt.Sprintf("User/Address/Create?userid=%s", user.ID()))
	contactsActions := pageactions.New(context, "UserContact", contactsOk)
	contactsActions.AddCreateWithPath(fmt.Sprintf("User/Contact/Create?userid=%s", user.ID()))

	if addressesOk {
		addressesActions.AddToggleWithPath("User/Address/ToggleActive")
		addressesActions.AddDeleteWithPath("User/Address/Delete")
		canUpdateAddress = context.HasUserRight("UpdateUserAddress")
		canDeleteAddress = context.HasUserRight("DeleteUserAddress")
	}

	if contactsOk {
		contactsActions.AddDeleteWithPath("User/Contact/Delete")
		canUpdateContact = context.HasUserRight("UpdateUserContact")
		canDeleteContact = context.HasUserRight("DeleteUserContact")
	}

	language, err := context.GetLanguage()
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	r.viewPageFactory.NewPage(context, language, user, addresses, addressesActions, contacts, contactsActions,
		canUpdateAddress, canDeleteAddress, canUpdateContact, canDeleteContact).Render()
}

// New returns a new instance of Route.
func New(userStore user.Store, userAddressStore address.Store, userContactStore contact.Store,
	viewPageFactory page.Factory) *Route {
	return &Route{
		userStore:        userStore,
		userAddressStore: userAddressStore,
		userContactStore: userContactStore,
		viewPageFactory:  viewPageFactory,
	}
}
