package search

import (
	"fmt"

	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/site"
	"github.com/juju/errors"
)

// Route processor.
type Route struct {
	siteStore site.Store
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

	sites, _, err := r.siteStore.Filter(context, language)
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}

	var firsthad bool
	for k := range sites {
		if firsthad {
			if _, err := context.WriteString("\n"); err != nil {
				context.RenderInternalServerError(errors.Trace(err))
				return
			}
		}
		if _, err := context.WriteString(fmt.Sprintf("%s\t", sites[k].ID())); err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
		if _, err := context.WriteString(r.siteStore.GetTranslatedName(sites[k], language.ID())); err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
		if !firsthad {
			firsthad = true
		}
	}
}

// New returns a new instance of Route.
func New(siteStore site.Store) *Route {
	return &Route{
		siteStore: siteStore,
	}
}
