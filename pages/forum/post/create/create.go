package create

import (
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/forum"
	"github.com/espal-digital-development/espal-module-core/pages/base"
)

var _ Factory = &Create{}
var _ Template = &Page{}

// Factory represents an object that serves new pages.
type Factory interface {
	NewPage(context contexts.Context, forum *forum.Forum) Template
}

// Create page service.
type Create struct{}

// NewPage generates a new instance of Page based on the given parameters.
func (c *Create) NewPage(context contexts.Context, forum *forum.Forum) Template {
	page := &Page{
		forum: forum,
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
	forum *forum.Forum
}

// Render the page writing to the context.
func (p *Page) Render() {
	base.WritePageTemplate(p.GetCoreContext(), p)
}

// New returns a new instance of Create.
func New() *Create {
	return &Create{}
}
