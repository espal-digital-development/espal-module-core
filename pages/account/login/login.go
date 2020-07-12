package login

import (
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-module-core/pages/base"
)

var _ Factory = &Login{}
var _ Template = &Page{}

// Factory represents an object that serves new pages.
type Factory interface {
	NewPage(context contexts.Context) Template
}

// Login page service.
type Login struct{}

// NewPage generates a new instance of Page based on the given parameters.
func (l *Login) NewPage(context contexts.Context) Template {
	page := &Page{}
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
}

// Render the page writing to the context.
func (p *Page) Render() {
	base.WritePageTemplate(p.GetCoreContext(), p)
}

// New returns a new instance of Login.
func New() *Login {
	return &Login{}
}
