package succeeded

import (
	"net/http"

	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	page "github.com/espal-digital-development/espal-module-core/pages/account/password/forgot/succeeded"
)

// Route processor.
type Route struct {
	succeededPageFactory page.Factory
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	if context.IsLoggedIn() {
		context.Redirect("/", http.StatusTemporaryRedirect)
		return
	}

	r.succeededPageFactory.NewPage(context).Render()
}

// New returns a new instance of Route.
func New(succeededPageFactory page.Factory) *Route {
	return &Route{
		succeededPageFactory: succeededPageFactory,
	}
}
