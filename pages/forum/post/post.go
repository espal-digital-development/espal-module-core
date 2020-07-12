package post

import (
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/forum"
	"github.com/espal-digital-development/espal-core/stores/user"
	"github.com/espal-digital-development/espal-core/template/renderer"
	"github.com/espal-digital-development/espal-module-core/pages/base"
)

var _ Factory = &Post{}
var _ Template = &Page{}

// Factory represents an object that serves new pages.
type Factory interface {
	NewPage(context contexts.Context, language contexts.Language, user *user.User, postEntity *forum.Post,
		replies []*forum.Post) Template
}

// Post page service.
type Post struct {
	rendererService renderer.Renderer
}

// NewPage generates a new instance of Page based on the given parameters.
func (p *Post) NewPage(context contexts.Context, language contexts.Language, user *user.User, postEntity *forum.Post,
	replies []*forum.Post) Template {
	page := &Page{
		language:        language,
		user:            user,
		post:            postEntity,
		replies:         replies,
		rendererService: p.rendererService,
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
	language        contexts.Language
	user            *user.User
	post            *forum.Post
	replies         []*forum.Post
	rendererService renderer.Renderer
}

// Render the page writing to the context.
func (p *Page) Render() {
	base.WritePageTemplate(p.GetCoreContext(), p)
}

// New returns a new instance of Post.
func New(rendererService renderer.Renderer) *Post {
	return &Post{
		rendererService: rendererService,
	}
}
