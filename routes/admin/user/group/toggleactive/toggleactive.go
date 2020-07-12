package toggleactive

import (
	"net/http"
	"strings"

	"github.com/espal-digital-development/espal-core/repositories/regularexpressions"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/user/group"
	"github.com/juju/errors"
)

// Route processor.
type Route struct {
	regularExpressionsRepository regularexpressions.Repository
	userGroupStore               group.Store
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	if !context.HasUserRightOrForbid("UpdateUserGroup") {
		return
	}
	ids := context.QueryValue("ids")
	if !r.regularExpressionsRepository.GetRouteIDs().MatchString(ids) {
		context.RenderBadRequest()
		return
	}
	if err := r.userGroupStore.ToggleActive(strings.Split(ids, ",")); err != nil {
		if err := context.SetFlashErrorMessage(context.Translate("activeTogglingHasFailed") + ": " +
			err.Error()); err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
	} else {
		if err := context.SetFlashSuccessMessage(context.Translate("activeTogglingWasSuccessful")); err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
	}
	context.Redirect(context.Referer(), http.StatusTemporaryRedirect)
}

// New returns a new instance of Route.
func New(regularExpressionsRepository regularexpressions.Repository, userGroupStore group.Store) *Route {
	return &Route{
		regularExpressionsRepository: regularExpressionsRepository,
		userGroupStore:               userGroupStore,
	}
}
