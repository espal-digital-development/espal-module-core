package delete

import (
	"net/http"

	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/forum"
	"github.com/juju/errors"
)

// Route processor.
type Route struct {
	forumStore forum.Store
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	if !context.HasUserRightOrForbid("DeleteForum") {
		return
	}
	id := context.QueryValue("id")
	if id == "" {
		context.RenderNotFound()
		return
	}

	forumID, ok, err := r.forumStore.GetForumIDForPostID(id)
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	if !ok {
		context.RenderNotFound()
		return
	}

	if err = r.forumStore.DeleteOneForumPostByID(id); err != nil {
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

	// TODO :: Also include showing it in the page(s)
	context.Redirect("/Forum?id="+forumID, http.StatusTemporaryRedirect)
}

// New returns a new instance of Route.
func New(forumStore forum.Store) *Route {
	return &Route{
		forumStore: forumStore,
	}
}
