package delete

import (
	"net/http"
	"strings"

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
	if !context.HasUserRightOrForbid("DeleteUser") {
		return
	}
	ids := context.QueryValue("ids")
	if !r.regularExpressionsRepository.GetRouteIDs().MatchString(ids) {
		context.RenderBadRequest()
		return
	}
	if err := r.userStore.Delete(strings.Split(ids, ",")); err != nil {
		if err := context.SetFlashErrorMessage(context.Translate("deletionHasFailed") + ": " + err.Error()); err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
	} else {
		if err := context.SetFlashSuccessMessage(context.Translate("deletionWasSuccessful")); err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
	}
	context.Redirect(context.Referer(), http.StatusTemporaryRedirect)
}

// New returns a new instance of Route.
func New(regularExpressionsRepository regularexpressions.Repository, userStore user.Store) *Route {
	return &Route{
		regularExpressionsRepository: regularExpressionsRepository,
		userStore:                    userStore,
	}
}
