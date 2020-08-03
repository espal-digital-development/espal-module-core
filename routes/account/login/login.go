package login

import (
	"github.com/espal-digital-development/espal-core/repositories/themes"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-module-core/forms/account/login"
	"github.com/juju/errors"
)

// Route processor.
type Route struct {
	loginFormValidator login.Factory
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	// TODO :: 777777 Fully implement this again based on the old view/route as a test
	viewData := themes.NewViewData()
	viewData.Set("test", "test text")
	if err := context.RenderTheme("login", viewData); err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
}

// New returns a new instance of Route.
func New(loginFormValidator login.Factory) *Route {
	return &Route{
		loginFormValidator: loginFormValidator,
	}
}
