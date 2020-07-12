package view

import (
	"github.com/espal-digital-development/espal-core/repositories/countries"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/site"
	"github.com/espal-digital-development/espal-module-core/pages/admin/base"
)

var _ Factory = &View{}
var _ Template = &Page{}

// Factory represents an object that serves new pages.
type Factory interface {
	NewPage(context contexts.Context, language contexts.Language, site *site.Site, translatedName string,
		siteLanguage contexts.Language, siteCountry countries.Data, currencies []string) Template
}

// View page service.
type View struct{}

// NewPage generates a new instance of Page based on the given parameters.
func (v *View) NewPage(context contexts.Context, language contexts.Language, site *site.Site, translatedName string,
	siteLanguage contexts.Language, siteCountry countries.Data, currencies []string) Template {
	page := &Page{
		language:       language,
		site:           site,
		translatedName: translatedName,
		siteLanguage:   siteLanguage,
		siteCountry:    siteCountry,
		currencies:     currencies,
	}
	page.SetCoreContext(context)
	return page
}

// Template represents a renderable page template object.
type Template interface {
	Render()
}

// Page contains and handles template logic.
type Page struct {
	base.Page
	language       contexts.Language
	site           *site.Site
	translatedName string
	siteLanguage   contexts.Language
	siteCountry    countries.Data
	currencies     []string
}

// Render the page writing to the context.
func (p *Page) Render() {
	base.WritePageTemplate(p.GetCoreContext(), p)
}

// New returns a new instance of View.
func New() *View {
	return &View{}
}
