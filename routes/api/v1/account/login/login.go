package login

import (
	"net/http"

	"github.com/espal-digital-development/espal-core/config"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/user"
)

// Route processor.
type Route struct {
	configService config.Config
	usersStore    user.Store
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	// TODO :: Deny overflow attacks and also check JWT token before allowing a new login request.
	email, err := context.FormValue("email")
	if err != nil {
		context.RenderInternalServerError(err)
		return
	}
	password, err := context.FormValue("password")
	if err != nil {
		context.RenderInternalServerError(err)
		return
	}
	if email == "" || password == "" {
		context.SetStatusCode(http.StatusUnauthorized)
		return
	}

	token, err := context.AuthorizeUserForJWT(email, password)
	if err != nil {
		context.RenderInternalServerError(err)
		return
	}

	context.WriteString(token)
}

// New returns a new instance of Route.
func New(configService config.Config, usersStore user.Store) *Route {
	return &Route{
		configService: configService,
		usersStore:    usersStore,
	}
}
