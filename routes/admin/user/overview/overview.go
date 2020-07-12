package overview

import (
	"github.com/espal-digital-development/espal-core/pageactions"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/user"
	page "github.com/espal-digital-development/espal-module-core/pages/admin/user/overview"
	"github.com/juju/errors"
)

// Route processor.
type Route struct {
	userStore           user.Store
	overviewPageFactory page.Factory
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	if !context.HasUserRightOrForbid("ReadUser") {
		return
	}

	users, filter, err := r.userStore.Filter(context)
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}

	var canUpdate bool
	var canDelete bool
	if filter.HasResults() {
		canUpdate = context.HasUserRight("UpdateUser")
		canDelete = context.HasUserRight("DeleteUser")
	}

	language, err := context.GetLanguage()
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	pageActions := pageactions.New(context, "User", len(users) > 0)
	pageActions.AddCreate()
	pageActions.AddToggle()
	pageActions.AddDelete()
	r.overviewPageFactory.NewPage(context, language, pageActions, filter, users, canUpdate, canDelete).Render()
}

// New returns a new instance of Route.
func New(userStore user.Store, overviewPageFactory page.Factory) *Route {
	return &Route{
		userStore:           userStore,
		overviewPageFactory: overviewPageFactory,
	}
}
