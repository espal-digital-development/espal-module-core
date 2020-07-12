package overview

import (
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/forum"
	page "github.com/espal-digital-development/espal-module-core/pages/forum/overview"
	"github.com/juju/errors"
)

// Route processor.
type Route struct {
	forumStore          forum.Store
	overviewPageFactory page.Factory
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	language, err := context.GetLanguage()
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}

	forums, _, err := r.forumStore.GetTopLevel(language)
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}

	r.overviewPageFactory.NewPage(context, forums, language).Render()
}

// New returns a new instance of Route.
func New(forumStore forum.Store, overviewPageFactory page.Factory) *Route {
	return &Route{
		forumStore:          forumStore,
		overviewPageFactory: overviewPageFactory,
	}
}
