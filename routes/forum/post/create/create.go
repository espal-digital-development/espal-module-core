package create

import (
	"github.com/espal-digital-development/espal-core/repositories/regularexpressions"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/forum"
	page "github.com/espal-digital-development/espal-module-core/pages/forum/post/create"
	"github.com/juju/errors"
)

// Route processor.
type Route struct {
	regularExpressionsRepository regularexpressions.Repository
	forumStore                   forum.Store
	createPageFactory            page.Factory
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	id := context.QueryValue("id")
	if !r.regularExpressionsRepository.GetRouteIDs().MatchString(id) {
		context.RenderBadRequest()
		return
	}
	language, err := context.GetLanguage()
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	forum, ok, err := r.forumStore.GetOneByID(id, language)
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	if !ok {
		context.RenderNotFound()
		return
	}
	// TODO :: Form, Post and Process
	r.createPageFactory.NewPage(context, forum).Render()
}

// New returns a new instance of Route.
func New(regularExpressionsRepository regularexpressions.Repository, forumStore forum.Store,
	createPageFactory page.Factory) *Route {
	return &Route{
		regularExpressionsRepository: regularExpressionsRepository,
		forumStore:                   forumStore,
		createPageFactory:            createPageFactory,
	}
}
