package userrightupdate

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
	id := context.QueryValue("id")
	if id == "" {
		context.RenderNotFound()
		return
	}
	if !r.regularExpressionsRepository.GetRouteIDs().MatchString(id) {
		context.RenderBadRequest()
		return
	}
	ids := context.QueryValue("ids")
	if !r.regularExpressionsRepository.GetRouteIDs().MatchString(ids) {
		context.RenderBadRequest()
		return
	}
	if err := r.userGroupStore.SetUserRights(id, strings.Split(ids, ",")); err != nil {
		if err := context.SetFlashErrorMessage(context.Translate("updateHasFailed") + ": " + err.Error()); err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
	} else {
		if err := context.SetFlashSuccessMessage(context.Translate("updateWasSuccessful")); err != nil {
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
