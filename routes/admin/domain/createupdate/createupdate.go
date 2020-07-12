package createupdate

import (
	"net/http"
	"strings"

	"github.com/espal-digital-development/espal-core/database/entitymutators"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/domain"
	"github.com/espal-digital-development/espal-core/validators/forms/admin/domain/createupdate"
	page "github.com/espal-digital-development/espal-module-core/pages/admin/domain/createupdate"
	"github.com/juju/errors"
)

// Route processor.
type Route struct {
	entityMutatorsFactory     entitymutators.Factory
	domainStore               domain.Store
	createUpdateFormValidator createupdate.Factory
	createUpdatePageFactory   page.Factory
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	id := context.QueryValue("id")

	if (id == "" && !context.HasUserRightOrForbid("CreateDomain")) ||
		(id != "" && !context.HasUserRightOrForbid("UpdateDomain")) {
		return
	}

	var err error
	var ok bool
	domain := &domain.Domain{}
	if id != "" {
		domain, ok, err = r.domainStore.GetOneByIDWithCreator(id)
		if err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
		if !ok {
			context.RenderNotFound()
			return
		}
	}

	language, err := context.GetLanguage()
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}

	form, err := r.createUpdateFormValidator.New(domain, language)
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	defer form.Close()
	isSubmitted, isValid, err := form.Submit(context)
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}

	if isSubmitted && isValid {
		siteID := form.FieldValue("site")
		languageID := form.FieldValueAsUint16("language")
		currencies := strings.Join(form.FieldValues("currencies"), ",")
		host := form.FieldValue("host")

		entityMutator := r.entityMutatorsFactory.NewMutation(domain, form, "Domain")
		entityMutator.SetBool("active", form.FieldValueAsBool("active"), domain.Active())
		entityMutator.SetString("siteID", &siteID, domain.SiteID())
		entityMutator.SetString("host", &host, domain.Host())
		entityMutator.SetNullableUint16("language", &languageID, domain.Language())
		entityMutator.SetString("currencies", &currencies, domain.Currencies())

		if err := entityMutator.Execute(context); err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}

		if redirectTo := entityMutator.RedirectURL(); redirectTo != "" {
			context.Redirect(context.AdminURL()+redirectTo, http.StatusTemporaryRedirect)
			return
		}

		domain, ok, err = r.domainStore.GetOneByIDWithCreator(entityMutator.GetInsertedOrUpdatedID())
		if err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
		if !ok {
			context.RenderNotFound()
			return
		}
	}

	r.createUpdatePageFactory.NewPage(context, domain, language, form.View(),
		context.GetAdminCreateUpdateTitle(id, "domain")).Render()
}

// New returns a new instance of Route.
func New(entityMutatorsFactory entitymutators.Factory, domainStore domain.Store,
	createUpdateFormValidator createupdate.Factory, createUpdatePageFactory page.Factory) *Route {
	return &Route{
		entityMutatorsFactory:     entityMutatorsFactory,
		domainStore:               domainStore,
		createUpdateFormValidator: createUpdateFormValidator,
		createUpdatePageFactory:   createUpdatePageFactory,
	}
}
