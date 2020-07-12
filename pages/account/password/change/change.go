package change

import (
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/template/renderer"
	"github.com/espal-digital-development/espal-module-core/pages/base"
)

// Factory represents an object that serves new pages.
type Factory interface {
	NewPage(context contexts.Context, form base.Form) Template
}

// Change page service.
type Change struct {
	rendererService renderer.Renderer
}

// NewPage generates a new instance of Page based on the given parameters.
func (c *Change) NewPage(context contexts.Context, form base.Form) Template {
	page := &Page{
		form:            form,
		rendererService: c.rendererService,
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
	form            base.Form
	rendererService renderer.Renderer
}

// Render the page writing to the context.
func (p *Page) Render() {
	base.WritePageTemplate(p.GetCoreContext(), p)
}

// New returns a new instance of Change.
func New(rendererService renderer.Renderer) *Change {
	return &Change{
		rendererService: rendererService,
	}
}
