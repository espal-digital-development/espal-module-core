package view

import (
	"strconv"
	"strings"

	"github.com/espal-digital-development/espal-core/repositories/countries"
	"github.com/espal-digital-development/espal-core/repositories/currencies"
	"github.com/espal-digital-development/espal-core/repositories/languages"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/site"
	page "github.com/espal-digital-development/espal-module-core/pages/admin/site/view"
	"github.com/juju/errors"
)

// Route processor.
type Route struct {
	languagesRepository  languages.Repository
	countriesRepository  countries.Repository
	currenciesRepository currencies.Repository
	siteStore            site.Store
	viewPageFactory      page.Factory
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	if !context.HasUserRightOrForbid("ReadSite") {
		return
	}

	id := context.QueryValue("id")
	if id == "" {
		context.RenderNotFound()
		return
	}

	site, ok, err := r.siteStore.GetOneByIDWithCreator(id)
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

	// TODO :: 7 Get this from the Site object itself
	var siteCurrencies []string
	if site.Currencies() != "" {
		currencyStringIds := strings.Split(site.Currencies(), ",")
		for k := range currencyStringIds {
			currencyID, err := strconv.ParseUint(currencyStringIds[k], 10, 64)
			if err != nil {
				context.RenderInternalServerError(errors.Trace(err))
				return
			}
			currency, err := r.currenciesRepository.ByID(uint(currencyID))
			if err != nil {
				context.RenderInternalServerError(errors.Trace(err))
				return
			}
			siteCurrencies = append(siteCurrencies, currency.Translate(language.ID()))
		}
	}

	var siteLanguage languages.Data
	if site.Language() != nil {
		siteLanguage, err = r.languagesRepository.ByID(*site.Language())
		if err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
	}
	var siteCountry countries.Data
	if site.Country() != nil {
		siteCountry, err = r.countriesRepository.ByID(*site.Country())
		if err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
	}
	r.viewPageFactory.NewPage(context, language, site, r.siteStore.GetTranslatedName(site, language.ID()),
		siteLanguage, siteCountry, siteCurrencies).Render()
}

// New returns a new instance of Route.
func New(languagesRepository languages.Repository, countriesRepository countries.Repository,
	currenciesRepository currencies.Repository, siteStore site.Store, viewPageFactory page.Factory) *Route {
	return &Route{
		languagesRepository:  languagesRepository,
		countriesRepository:  countriesRepository,
		currenciesRepository: currenciesRepository,
		siteStore:            siteStore,
		viewPageFactory:      viewPageFactory,
	}
}
