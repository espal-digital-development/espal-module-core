package login

import (
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/validators/forms/account/login"
	page "github.com/espal-digital-development/espal-module-core/pages/account/login"
)

// Route processor.
type Route struct {
	loginFormValidator login.Factory
	loginPageFactory   page.Factory
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	r.loginPageFactory.NewPage(context).Render()
}

// New returns a new instance of Route.
func New(loginFormValidator login.Factory, loginPageFactory page.Factory) *Route {
	return &Route{
		loginFormValidator: loginFormValidator,
		loginPageFactory:   loginPageFactory,
	}
}
