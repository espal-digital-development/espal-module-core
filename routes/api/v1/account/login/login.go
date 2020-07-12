package login

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/user"
)

type JWTToken struct {
	jwt.StandardClaims
	UserID string
}

// Route processor.
type Route struct {
	usersStore user.Store
}

// TODO :: 777777 Make this a configService setting.
const tokenPassword = "42e1d1a0b8a66670a2a748a327dfffa5"

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	context.SetHeader("Access-Control-Allow-Origin", "*")

	email, err := context.FormValue("email")
	if err != nil {
		context.RenderInternalServerError(err)
		return
	}
	user, ok, err := r.usersStore.GetOneByEmail(email)
	if err != nil {
		context.RenderInternalServerError(err)
		return
	}
	if !ok {
		context.SetStatusCode(http.StatusUnauthorized)
		return
	}

	// TODO :: Validate password, but just for testing purposes now

	// TODO :: Must use Ed25519? But is not supported?
	tk := &JWTToken{UserID: user.ID()}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS512"), tk)
	tokenString, err := token.SignedString([]byte(tokenPassword))
	if err != nil {
		context.RenderInternalServerError(err)
		return
	}

	// context.SetContentType("text/plain")
	context.SetContentType("espal-x")

	if _, err := context.WriteString(tokenString); err != nil {
		context.SetStatusCode(http.StatusBadRequest)
		return
	}

}

// New returns a new instance of Route.
func New(usersStore user.Store) *Route {
	return &Route{
		usersStore: usersStore,
	}
}
