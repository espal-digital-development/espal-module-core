package register

import (
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-module-core/pages/base"
)

var _ Factory = &Register{}
var _ Template = &Page{}

// Factory represents an object that serves new pages.
type Factory interface {
	NewPage(context contexts.Context, form base.Form) Template
}

// Register page service.
type Register struct{}

// NewPage generates a new instance of Page based on the given parameters.
func (r *Register) NewPage(context contexts.Context, form base.Form) Template {
	page := &Page{
		form: form,
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
	form base.Form
}

// Render the page writing to the context.
func (p *Page) Render() {
	base.WritePageTemplate(p.GetCoreContext(), p)
}

// New returns a new instance of Register.
func New() *Register {
	return &Register{}
}
