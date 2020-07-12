package dashboard

import (
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-module-core/pages/admin/dashboard"
)

// Route processor.
type Route struct {
	dashboardPageFactory dashboard.Factory
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	if !context.HasAdminAccess() {
		context.RenderUnauthorized()
		return
	}

	r.dashboardPageFactory.NewPage(context).Render()
}

// New returns a new instance of Route.
func New(dashboardPageFactory dashboard.Factory) *Route {
	return &Route{
		dashboardPageFactory: dashboardPageFactory,
	}
}
