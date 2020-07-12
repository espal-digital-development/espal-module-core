package createupdate

import (
	"net/http"
	"strings"

	"github.com/espal-digital-development/espal-core/database/entitymutators"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/site"
	"github.com/espal-digital-development/espal-core/validators/forms/admin/site/createupdate"
	page "github.com/espal-digital-development/espal-module-core/pages/admin/site/createupdate"
	"github.com/juju/errors"
)

// Route processor.
type Route struct {
	entityMutatorsFactory     entitymutators.Factory
	siteStore                 site.Store
	createUpdateFormValidator createupdate.Factory
	createUpdatePageFactory   page.Factory
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	var err error
	id := context.QueryValue("id")

	if (id == "" && !context.HasUserRightOrForbid("CreateSite")) ||
		(id != "" && !context.HasUserRightOrForbid("UpdateSite")) {
		return
	}

	var ok bool
	site := &site.Site{}
	if id != "" {
		site, ok, err = r.siteStore.GetOneByIDWithCreator(id)
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

	form, err := r.createUpdateFormValidator.New(site, language)
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
		countryID := form.FieldValueAsUint16("country")
		languageID := form.FieldValueAsUint16("language")
		currencies := strings.Join(form.FieldValues("currencies"), ",")

		entityMutator := r.entityMutatorsFactory.NewMutation(site, form, "Site")
		entityMutator.SetBool("online", form.FieldValueAsBool("online"), site.Online())
		entityMutator.SetNullableUint16("language", &languageID, site.Language())
		entityMutator.SetNullableUint16("country", &countryID, site.Country())
		entityMutator.SetString("currencies", &currencies, site.Currencies())

		// TODO :: 7 Translations maybe through a new FormFieldType?

		if err := entityMutator.Execute(context); err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}

		if redirectTo := entityMutator.RedirectURL(); redirectTo != "" {
			context.Redirect(context.AdminURL()+redirectTo, http.StatusTemporaryRedirect)
			return
		}

		site, ok, err = r.siteStore.GetOneByIDWithCreator(entityMutator.GetInsertedOrUpdatedID())
		if err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
		if !ok {
			context.RenderNotFound()
			return
		}
	}

	// TODO :: 77 Creation succeeds but still shows `Validation token invalid/expired`
	r.createUpdatePageFactory.NewPage(context, site, language, form.View(),
		context.GetAdminCreateUpdateTitle(id, "site")).Render()
}

// New returns a new instance of Route.
func New(entityMutatorsFactory entitymutators.Factory, siteStore site.Store,
	createUpdateFormValidator createupdate.Factory, createUpdatePageFactory page.Factory) *Route {
	return &Route{
		entityMutatorsFactory:     entityMutatorsFactory,
		siteStore:                 siteStore,
		createUpdateFormValidator: createUpdateFormValidator,
		createUpdatePageFactory:   createUpdatePageFactory,
	}
}
