package post

import (
	"github.com/espal-digital-development/espal-core/repositories/regularexpressions"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/forum"
	"github.com/espal-digital-development/espal-core/stores/user"
	page "github.com/espal-digital-development/espal-module-core/pages/forum/post"
	"github.com/juju/errors"
)

// Route processor.
type Route struct {
	regularExpressionsRepository regularexpressions.Repository
	forumStore                   forum.Store
	postPageFactory              page.Factory
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	language, err := context.GetLanguage()
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	id := context.QueryValue("id")
	if !r.regularExpressionsRepository.GetRouteIDs().MatchString(id) {
		context.RenderBadRequest()
		return
	}

	forumPost, ok, err := r.forumStore.GetOnePostByID(id)
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	if !ok {
		context.RenderNotFound()
		return
	}

	replies, _, err := r.forumStore.GetPostReplies(id)
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}

	user := &user.User{}
	if context.IsLoggedIn() {
		user, ok, err = context.GetUser()
		if err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
		if !ok {
			context.RenderBadRequest()
			return
		}
	}

	r.postPageFactory.NewPage(context, language, user, forumPost, replies).Render()
}

// New returns a new instance of Route.
func New(regularExpressionsRepository regularexpressions.Repository, forumStore forum.Store,
	postPageFactory page.Factory) *Route {
	return &Route{
		regularExpressionsRepository: regularExpressionsRepository,
		forumStore:                   forumStore,
		postPageFactory:              postPageFactory,
	}
}
