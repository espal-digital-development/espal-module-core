package spa

import (
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	page "github.com/espal-digital-development/espal-module-core/pages/spa"
)

// Route processor.
type Route struct {
	spaPageFactory page.Factory
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	r.spaPageFactory.NewPage(context).Render()
}

// New returns a new instance of Route.
func New(spaPageFactory page.Factory) *Route {
	return &Route{
		spaPageFactory: spaPageFactory,
	}
}
