package global

import (
	"encoding/json"

	"github.com/espal-digital-development/espal-core/repositories/languages"
	"github.com/espal-digital-development/espal-core/repositories/userrights"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/domain"
)

// Route processor.
type Route struct {
	languagesRepository  languages.Repository
	userRightsRepository userrights.Repository
	domainStore          domain.Store
}

// nolint:funlen
// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	// TODO :: 777777 Caching and the notifier update
	site := context.GetSite()
	domain := context.GetDomain()

	repoDomains, ok, err := r.domainStore.AllForSiteID(site.ID())
	var domains []*domainJSON
	if err != nil {
		context.RenderInternalServerError(err)
		return
	}
	if ok {
		domains = make([]*domainJSON, len(repoDomains))
		for k := range repoDomains {
			domains[k] = &domainJSON{
				ID:       repoDomains[k].ID(),
				Host:     repoDomains[k].HostWithProtocol(),
				Language: repoDomains[k].Language(),
			}
		}
	}

	repoLanguages := r.languagesRepository.All()
	languages := []*languageJSON{}
	for id := range repoLanguages {
		var alternativeEnglishName *string
		alternativeEnglishNameData := repoLanguages[id].AlternativeEnglishName()
		if alternativeEnglishNameData != "" {
			alternativeEnglishName = &alternativeEnglishNameData
		}
		languages = append(languages, &languageJSON{
			ID:                     repoLanguages[id].ID(),
			EnglishName:            repoLanguages[id].EnglishName(),
			AlternativeEnglishName: alternativeEnglishName,
		})
	}

	storeObjects := []*storeObject{
		// TODO :: 777777 Auto-generated all these entries with gogenerate
		{
			Name: "User",
			Fields: []*storeObjectField{
				{
					Name: "ID",
					Type: "String",
				},
			},
		},
	}

	data := &globalJSON{
		Site: &siteJSON{
			ID:       site.ID(),
			Language: site.Language(),
		},
		Domain: &domainJSON{
			ID:       domain.ID(),
			Host:     domain.HostWithProtocol(),
			Language: domain.Language(),
		},
		Domains:      domains,
		Languages:    languages,
		UserRights:   r.userRightsRepository.AllByCode(),
		StoreObjects: storeObjects,
	}

	jsonBytes, err := json.Marshal(&data)
	if err != nil {
		context.RenderInternalServerError(err)
		return
	}
	context.SetContentType("application/json")
	context.Write(jsonBytes)
}

// New returns a new instance of Route.
func New(languagesRepository languages.Repository, userRightsRepository userrights.Repository,
	domainStore domain.Store) *Route {
	return &Route{
		languagesRepository:  languagesRepository,
		userRightsRepository: userRightsRepository,
		domainStore:          domainStore,
	}
}
