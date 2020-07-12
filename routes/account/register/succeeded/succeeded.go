package succeeded

import (
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-module-core/pages/account/register/succeeded"
)

// Route processor.
type Route struct {
	succeededPageFactory succeeded.Factory
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	r.succeededPageFactory.NewPage(context).Render()
}

// New returns a new instance of Route.
func New(succeededPageFactory succeeded.Factory) *Route {
	return &Route{
		succeededPageFactory: succeededPageFactory,
	}
}
