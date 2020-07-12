package forum

import (
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	forumstore "github.com/espal-digital-development/espal-core/stores/forum"
	"github.com/espal-digital-development/espal-core/template/renderer"
	"github.com/espal-digital-development/espal-module-core/pages/base"
)

var _ Factory = &Forum{}
var _ Template = &Page{}

// Factory represents an object that serves new pages.
type Factory interface {
	NewPage(context contexts.Context, language contexts.Language, forumEntity *forumstore.Forum,
		posts []*forumstore.Post, forums []*forumstore.Forum) Template
}

// Forum page service.
type Forum struct {
	rendererService renderer.Renderer
}

// NewPage generates a new instance of Page based on the given parameters.
func (f *Forum) NewPage(context contexts.Context, language contexts.Language, forumEntity *forumstore.Forum,
	posts []*forumstore.Post, forums []*forumstore.Forum) Template {
	page := &Page{
		language:        language,
		forum:           forumEntity,
		posts:           posts,
		forums:          forums,
		rendererService: f.rendererService,
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
	forum           *forumstore.Forum
	posts           []*forumstore.Post
	forums          []*forumstore.Forum
	rendererService renderer.Renderer
}

// Render the page writing to the context.
func (p *Page) Render() {
	base.WritePageTemplate(p.GetCoreContext(), p)
}

// New returns a new instance of Forum.
func New(rendererService renderer.Renderer) *Forum {
	return &Forum{
		rendererService: rendererService,
	}
}
