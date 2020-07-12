package overview

import (
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/forum"
	"github.com/espal-digital-development/espal-core/template/renderer"
	"github.com/espal-digital-development/espal-module-core/pages/base"
)

var _ Factory = &Overview{}
var _ Template = &Page{}

// Factory represents an object that serves new pages.
type Factory interface {
	NewPage(context contexts.Context, forums []*forum.Forum, language contexts.Language) Template
}

// Overview page service.
type Overview struct {
	rendererService renderer.Renderer
}

// NewPage generates a new instance of Page based on the given parameters.
func (o *Overview) NewPage(context contexts.Context, forums []*forum.Forum, language contexts.Language) Template {
	page := &Page{
		forums:          forums,
		language:        language,
		rendererService: o.rendererService,
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
	forums          []*forum.Forum
	language        contexts.Language
	rendererService renderer.Renderer
}

// Render the page writing to the context.
func (p *Page) Render() {
	base.WritePageTemplate(p.GetCoreContext(), p)
}

// New returns a new instance of Overview.
func New(rendererService renderer.Renderer) *Overview {
	return &Overview{
		rendererService: rendererService,
	}
}
