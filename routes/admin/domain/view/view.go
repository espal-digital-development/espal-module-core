package view

import (
	"github.com/espal-digital-development/espal-core/repositories/languages"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/domain"
	page "github.com/espal-digital-development/espal-module-core/pages/admin/domain/view"
	"github.com/juju/errors"
)

// Route processor.
type Route struct {
	languagesRepository languages.Repository
	domainStore         domain.Store
	viewPageFactory     page.Factory
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	if !context.HasUserRightOrForbid("ReadDomain") {
		return
	}

	id := context.QueryValue("id")
	if len(id) == 0 {
		context.RenderNotFound()
		return
	}

	domain, ok, err := r.domainStore.GetOneByIDWithCreator(id)
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	if !ok {
		context.RenderNotFound()
		return
	}

	language, err := context.GetLanguage()
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}

	var domainLanguage languages.Data
	if domain.Language() != nil {
		domainLanguage, err = r.languagesRepository.ByID(*domain.Language())
		if err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
	}
	r.viewPageFactory.NewPage(context, language, domain, domainLanguage).Render()
}

// New returns a new instance of Route.
func New(languagesRepository languages.Repository, domainStore domain.Store, viewPageFactory page.Factory) *Route {
	return &Route{
		languagesRepository: languagesRepository,
		domainStore:         domainStore,
		viewPageFactory:     viewPageFactory,
	}
}
