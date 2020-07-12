package spa

import (
	"github.com/espal-digital-development/espal-core/pages/base"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
)

var _ Factory = &SPA{}
var _ Template = &Page{}

// Factory represents an object that serves new pages.
type Factory interface {
	NewPage(context contexts.Context) Template
}

// SPA page service.
type SPA struct{}

// NewPage generates a new instance of Page based on the given parameters.
func (c *SPA) NewPage(context contexts.Context) Template {
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

// New returns a new instance of SPA.
func New() *SPA {
	return &SPA{}
}
