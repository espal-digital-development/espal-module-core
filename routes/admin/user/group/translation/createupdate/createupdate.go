package createupdate

import (
	"github.com/espal-digital-development/espal-core/repositories/regularexpressions"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/user/group"
)

// Route processor.
type Route struct {
	regularExpressionsRepository regularexpressions.Repository
	userGroupStore               group.Store
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	// TODO :: 77 Implement this (is being called from user/group/view.qtpl)
	// if !context.HasUserRightOrForbid("UpdateUserGroup") {
	// 	return
	// }

	// ids := context.QueryValue("ids")
	// if !r.regularExpressionsRepository.GetRouteIDs().MatchString(ids) {
	// 	context.RenderBadRequest()
	// 	return
	// }

	// idString := context.QueryValue("id")
	// if len(idString) == 0 {
	// 	context.RenderNotFound()
	// 	return
	// }
	// if !r.regularExpressionsRepository.GetRouteIDs().MatchString(idString) {
	// 	context.RenderBadRequest()
	// 	return
	// }

	// admin.WritePageTemplate(ctx, &admin.UserGroupTranslationCreateUpdatePage{
	// 	BasePage:     admin.BasePage{Route: routeCtx},
	// 	UserGroupTranslation:         userGroupTranslation,
	// 	Language:     language,
	// 	Form:         form.LoadViewData(),
	// 	DisplayTitle: context.GetAdminCreateUpdateTitle(id, "user"),
	// })
}

// New returns a new instance of Route.
func New(regularExpressionsRepository regularexpressions.Repository, userGroupStore group.Store) *Route {
	return &Route{
		regularExpressionsRepository: regularExpressionsRepository,
		userGroupStore:               userGroupStore,
	}
}
