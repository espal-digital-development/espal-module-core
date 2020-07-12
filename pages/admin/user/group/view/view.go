package view

import (
	"github.com/espal-digital-development/espal-core/pageactions"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/user/group"
	"github.com/espal-digital-development/espal-core/template/renderer"
	"github.com/espal-digital-development/espal-module-core/pages/admin/base"
)

var _ Factory = &View{}
var _ Template = &Page{}

// Factory represents an object that serves new pages.
type Factory interface {
	NewPage(context contexts.Context, language contexts.Language, userGroup *group.Group, userRights map[uint16]string,
		userRightsOrder []uint16, userGroupUserRights map[uint16]bool, userRightsActions pageactions.Actions,
		translations []*group.Translation, translationsActions pageactions.Actions, canUpdate bool,
		canDelete bool) Template
}

// View page service.
type View struct {
	rendererService renderer.Renderer
}

// NewPage generates a new instance of Page based on the given parameters.
func (v *View) NewPage(context contexts.Context, language contexts.Language, userGroup *group.Group,
	userRights map[uint16]string, userRightsOrder []uint16, userGroupUserRights map[uint16]bool,
	userRightsActions pageactions.Actions, translations []*group.Translation,
	translationsActions pageactions.Actions, canUpdate bool, canDelete bool) Template {
	page := &Page{
		language:            language,
		userGroup:           userGroup,
		userRights:          userRights,
		userRightsOrder:     userRightsOrder,
		userGroupUserRights: userGroupUserRights,
		userRightsActions:   userRightsActions,
		translations:        translations,
		translationsActions: translationsActions,
		canUpdate:           canUpdate,
		canDelete:           canDelete,
		rendererService:     v.rendererService,
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
	language            contexts.Language
	userGroup           *group.Group
	userRights          map[uint16]string
	userRightsOrder     []uint16
	userGroupUserRights map[uint16]bool
	userRightsActions   pageactions.Actions
	translations        []*group.Translation
	translationsActions pageactions.Actions
	canUpdate           bool
	canDelete           bool
	rendererService     renderer.Renderer
}

// Render the page writing to the context.
func (p *Page) Render() {
	base.WritePageTemplate(p.GetCoreContext(), p)
}

// New returns a new instance of View.
func New(rendererService renderer.Renderer) *View {
	return &View{
		rendererService: rendererService,
	}
}
