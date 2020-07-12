package logout

import (
	"net/http"

	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/juju/errors"
)

// Route processor.
type Route struct{}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	if !context.IsLoggedIn() {
		context.Redirect("/", http.StatusTemporaryRedirect)
		return
	}
	if err := context.Logout(); err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	context.Redirect("/", http.StatusTemporaryRedirect)
}

// New returns a new instance of Route.
func New() *Route {
	return &Route{}
}
