package overview

import (
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	page "github.com/espal-digital-development/espal-module-core/pages/account/overview"
)

// Route processor.
type Route struct {
	overviewPageFactory page.Factory
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	r.overviewPageFactory.NewPage(context).Render()
}

// New returns a new instance of Route.
func New(overviewPageFactory page.Factory) *Route {
	return &Route{
		overviewPageFactory: overviewPageFactory,
	}
}
