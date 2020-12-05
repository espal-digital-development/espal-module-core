package core

import (
	"path"
	"runtime"

	"github.com/espal-digital-development/espal-core/modules"
	"github.com/espal-digital-development/espal-core/modules/assets"
	"github.com/espal-digital-development/espal-core/modules/meta"
	"github.com/espal-digital-development/espal-core/modules/repositories"
	"github.com/espal-digital-development/espal-core/modules/routes"
	"github.com/espal-digital-development/espal-core/modules/themes"
	"github.com/espal-digital-development/espal-core/modules/translations"
	themesRepository "github.com/espal-digital-development/espal-core/repositories/themes"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	loginForm "github.com/espal-digital-development/espal-module-core/forms/account/login"
	"github.com/espal-digital-development/espal-module-core/pages/spa"
	loginRoute "github.com/espal-digital-development/espal-module-core/routes/account/login"
	apiAccountLoginRoute "github.com/espal-digital-development/espal-module-core/routes/api/v1/account/login"
	apiAccountOverviewRoute "github.com/espal-digital-development/espal-module-core/routes/api/v1/account/overview"
	"github.com/espal-digital-development/espal-module-core/routes/api/v1/global"
	spaRoute "github.com/espal-digital-development/espal-module-core/routes/spa"
	"github.com/espal-digital-development/espal-module-core/themes/base/login"
	"github.com/juju/errors"
)

var errResolveModulePath = errors.New("failed to resolve module path")

// TODO :: 777777
// - How to hook into existing functionality like Slugs and other Database/Repository functionality
//   - How to get functionality báck into the modules? Some kind of reverse registration injection with interface{}'s?
// - CompatibilityDefintion should describe what versions of the core app works with and
//   whether it colides with other functionality being present in the system through other modules
// - Modules should be able to talk to each other and affect load order

// TODO :: 777777 Register forms and other routes/views (migrate all pages into the new Theme Views)

// New returns a new instance of Module.
// nolint:funlen
func New() (*modules.Module, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, errors.Trace(errResolveModulePath)
	}
	modulePath := path.Dir(filename)
	config := &modules.Config{}
	var err error

	config.MetaDefinition, err = meta.New(&meta.Config{
		UniqueIdentifier:             "com.espal.core",
		Version:                      "0.0.1",
		MinimumCompatibleCoreVersion: "0.0.1",
		MaximumCompatibleCoreVersion: "",
		Name:                         "Espal Core",
		Author:                       "Espal Digital Development",
		Contact:                      "https://github.com/espal-digital-development",
	})
	if err != nil {
		return nil, errors.Trace(err)
	}

	config.PreGetAssetsCallback = func(m modules.Modular) error {
		assets, err := assets.New(&assets.Config{
			PublicRootFilesPath: modulePath + "/app/assets/files/root",
			ImagesPath:          modulePath + "/app/assets/images",
			StylesheetsPath:     modulePath + "/app/assets/css",
			JavaScriptPath:      modulePath + "/app/assets/js",
		})
		if err != nil {
			return errors.Trace(err)
		}
		m.SetAssets(assets)
		return nil
	}

	config.PreGetTranslationsCallback = func(m modules.Modular) error {
		translations, err := translations.New(&translations.Config{
			Path: modulePath + "/app/translations",
		})
		if err != nil {
			return errors.Trace(err)
		}
		m.SetTranslations(translations)
		return nil
	}

	config.PreGetThemesCallback = func(m modules.Modular) error {
		loginView := login.New()

		themes, err := themes.New(&themes.Config{
			Views: map[string][]themesRepository.Viewable{
				"base": {
					loginView,
				},
			},
		})
		if err != nil {
			return errors.Trace(err)
		}
		m.SetThemes(themes)
		return nil
	}

	config.PreGetRepositoriesCallback = func(m modules.Modular) error {
		repostories, err := repositories.New()
		if err != nil {
			return errors.Trace(err)
		}
		m.SetRepositories(repostories)
		return nil
	}

	config.PreGetRoutesCallback = func(m modules.Modular) error {
		repos, err := m.GetRepositories()
		if err != nil {
			return errors.Trace(err)
		}

		routes, err := routes.New(&routes.Config{
			Entries: map[string]routes.Handler{
				"/": spaRoute.New(spa.New()),

				"/API/V1/Global": global.New(repos.Languages(), repos.UserRights(), m.GetStores().Domain()),

				"/API/V1/Account":                         apiAccountOverviewRoute.New(m.GetStores().User()),
				"/API/V1/Login":                           apiAccountLoginRoute.New(m.GetStores().User()),
				"/API/V1/Account/Register":                &apiEndPointNotImplemented{},
				"/API/V1/Account/ForgotPassword":          &apiEndPointNotImplemented{},
				"/API/V1/Account/ForgotPasswordSucceeded": &apiEndPointNotImplemented{},
				"/API/V1/Account/PasswordRecovery":        &apiEndPointNotImplemented{},

				"/Login": loginRoute.New(loginForm.New(m.GetValidatorsFactory(), m.GetStores().User())),
			},
		})
		if err != nil {
			return errors.Trace(err)
		}
		m.SetRoutes(routes)
		return nil
	}

	m, err := modules.New(config)
	return m, errors.Trace(err)
}

type apiEndPointNotImplemented struct{}

func (a *apiEndPointNotImplemented) Handle(context contexts.Context) {
	context.WriteString("This endpoint is not implemented yet.")
}
