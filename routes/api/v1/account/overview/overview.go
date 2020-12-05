package overview

import (
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
)

// Route processor.
type Route struct{}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	if !context.IsJWTAuthorized() {
		return
	}
}

// New returns a new instance of Route.
func New() *Route {
	return &Route{}
}
