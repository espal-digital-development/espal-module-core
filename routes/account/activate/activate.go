package activate

import (
	"net/http"

	"github.com/espal-digital-development/espal-core/repositories/regularexpressions"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/user"
	"github.com/juju/errors"
)

// Route processor.
type Route struct {
	regularExpressionsRepository regularexpressions.Repository
	userStore                    user.Store
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	if context.IsLoggedIn() {
		context.Redirect("/", http.StatusTemporaryRedirect)
		return
	}

	hash := context.QueryValue("h")
	if hash == "" || !r.regularExpressionsRepository.GetActivateAccounthash().MatchString(hash) {
		context.RenderBadRequest()
		return
	}

	id, ok, err := r.userStore.GetOneIDForActivationHash(hash)
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	if !ok {
		// TODO :: 7 Support email and/or internal Support requests/chat module?
		context.SetStatusCode(http.StatusBadRequest)
		context.RenderNon200Custom(context.Translate("unknownActivationHash"),
			context.Translate("unknownActivationHashMessage"))
		return
	}

	if err := r.userStore.Activate(id); err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}

	// TODO :: Setting-bool to mail a welcome mail?

	// TODO :: Redirect to a new Activated page instead of redirecting to Login directly
	context.Redirect("/Login", http.StatusTemporaryRedirect)
}

// New returns a new instance of Route.
func New(regularExpressionsRepository regularexpressions.Repository, userStore user.Store) *Route {
	return &Route{
		regularExpressionsRepository: regularExpressionsRepository,
		userStore:                    userStore,
	}
}
