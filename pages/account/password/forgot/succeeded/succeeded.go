package succeeded

import (
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-module-core/pages/base"
)

var _ Factory = &Succeeded{}
var _ Template = &Page{}

// Factory represents an object that serves new pages.
type Factory interface {
	NewPage(context contexts.Context) Template
}

// Succeeded page service.
type Succeeded struct{}

// NewPage generates a new instance of Page based on the given parameters.
func (s *Succeeded) NewPage(context contexts.Context) Template {
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

// New returns a new instance of Succeeded.
func New() *Succeeded {
	return &Succeeded{}
}
