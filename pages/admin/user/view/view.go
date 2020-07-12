package view

import (
	"github.com/espal-digital-development/espal-core/pageactions"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/user"
	"github.com/espal-digital-development/espal-core/stores/user/address"
	"github.com/espal-digital-development/espal-core/stores/user/contact"
	"github.com/espal-digital-development/espal-core/template/renderer"
	"github.com/espal-digital-development/espal-module-core/pages/admin/base"
)

var _ Factory = &View{}
var _ Template = &Page{}

// Factory represents an object that serves new pages.
type Factory interface {
	NewPage(context contexts.Context, language contexts.Language, user *user.User, addresses []*address.Address,
		addressesActions pageactions.Actions, contacts []*contact.Contact, contactsActions pageactions.Actions,
		canUpdateAddress bool, canDeleteAddress bool, canUpdateContact bool, canDeleteContact bool) Template
}

// View page service.
type View struct {
	rendererService  renderer.Renderer
	userStore        user.Store
	userContactStore contact.Store
}

// NewPage generates a new instance of Page based on the given parameters.
func (v *View) NewPage(context contexts.Context, language contexts.Language, user *user.User,
	addresses []*address.Address, addressesActions pageactions.Actions, contacts []*contact.Contact,
	contactsActions pageactions.Actions, canUpdateAddress bool, canDeleteAddress bool, canUpdateContact bool,
	canDeleteContact bool) Template {
	page := &Page{
		language:         language,
		user:             user,
		addresses:        addresses,
		addressesActions: addressesActions,
		contacts:         contacts,
		contactsActions:  contactsActions,
		canUpdateAddress: canUpdateAddress,
		canDeleteAddress: canDeleteAddress,
		canUpdateContact: canUpdateContact,
		canDeleteContact: canDeleteContact,
		rendererService:  v.rendererService,
		userStore:        v.userStore,
		userContactStore: v.userContactStore,
	}
	page.SetCoreContext(context)
	return page
}

// Template represents a renderable page template object.
type Template interface {
	Render()
}

// Page contains and handles template logic.
type Page struct {
	base.Page
	language         contexts.Language
	user             *user.User
	addresses        []*address.Address
	addressesActions pageactions.Actions
	contacts         []*contact.Contact
	contactsActions  pageactions.Actions
	canUpdateAddress bool
	canDeleteAddress bool
	canUpdateContact bool
	canDeleteContact bool
	rendererService  renderer.Renderer
	userStore        user.Store
	userContactStore contact.Store
}

// Render the page writing to the context.
func (p *Page) Render() {
	base.WritePageTemplate(p.GetCoreContext(), p)
}

// New returns a new instance of View.
func New(rendererService renderer.Renderer, userStore user.Store, userContactStore contact.Store) *View {
	return &View{
		rendererService:  rendererService,
		userStore:        userStore,
		userContactStore: userContactStore,
	}
}
