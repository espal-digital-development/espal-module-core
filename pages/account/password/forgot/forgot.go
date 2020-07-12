package forgot

import (
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-module-core/pages/base"
)

var _ Factory = &Forgot{}
var _ Template = &Page{}

// Factory represents an object that serves new pages.
type Factory interface {
	NewPage(context contexts.Context, form base.Form) Template
}

// Forgot page service.
type Forgot struct{}

// NewPage generates a new instance of Page based on the given parameters.
func (f *Forgot) NewPage(context contexts.Context, form base.Form) Template {
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

// New returns a new instance of Forgot.
func New() *Forgot {
	return &Forgot{}
}
