package overview

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/user"
)

const tokenHeaderChunks = 2

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
	tokenHeader := context.GetHeader("Authorization")
	if tokenHeader == "" {
		context.SetStatusCode(http.StatusForbidden)
		return
	}
	splitted := strings.Split(tokenHeader, " ")
	if len(splitted) != tokenHeaderChunks {
		// "Invalid/Malformed auth token"
		context.SetStatusCode(http.StatusForbidden)
		return
	}
	tokenPart := splitted[1]
	tk := &JWTToken{}
	token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenPassword), nil
	})
	if err != nil {
		// "Malformed authentication token"
		context.SetStatusCode(http.StatusForbidden)
		return
	}
	if !token.Valid {
		context.SetStatusCode(http.StatusForbidden)
		return
	}

	// context.SetContentType("text/plain")
	context.SetContentType("espal-x")

	if _, err := context.WriteString("Your token is valid!"); err != nil {
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
