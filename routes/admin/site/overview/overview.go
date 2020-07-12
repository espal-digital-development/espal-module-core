package overview

import (
	"github.com/espal-digital-development/espal-core/pageactions"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/site"
	page "github.com/espal-digital-development/espal-module-core/pages/admin/site/overview"
	"github.com/juju/errors"
)

// Route processor.
type Route struct {
	siteStore           site.Store
	overviewPageFactory page.Factory
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	if !context.HasUserRightOrForbid("ReadSite") {
		return
	}
	language, err := context.GetLanguage()
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}

	sites, filter, err := r.siteStore.Filter(context, language)
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}

	var canUpdate bool
	var canDelete bool
	if filter.HasResults() {
		canUpdate = context.HasUserRight("UpdateSite")
		canDelete = context.HasUserRight("DeleteSite")
	}
	pageActions := pageactions.New(context, "Site", len(sites) > 0)
	pageActions.AddCreate()
	pageActions.AddToggle()
	pageActions.AddDelete()
	r.overviewPageFactory.NewPage(context, language, pageActions, filter, sites, canUpdate, canDelete).Render()
}

// New returns a new instance of Route.
func New(siteStore site.Store, overviewPageFactory page.Factory) *Route {
	return &Route{
		siteStore:           siteStore,
		overviewPageFactory: overviewPageFactory,
	}
}
