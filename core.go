package core

import (
	"net/http"
	"path"
	"path/filepath"
	"runtime"

	"github.com/espal-digital-development/espal-core/modules"
	"github.com/espal-digital-development/espal-core/modules/assets"
	"github.com/espal-digital-development/espal-core/modules/meta"
	"github.com/espal-digital-development/espal-core/modules/routes"
	"github.com/espal-digital-development/espal-core/modules/translations"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-module-core/pages/catalog"
	"github.com/espal-digital-development/espal-module-core/pages/root"
	"github.com/espal-digital-development/espal-module-core/pages/spa"
	apiAccountLoginRoute "github.com/espal-digital-development/espal-module-core/routes/api/v1/account/login"
	apiAccountOverviewRoute "github.com/espal-digital-development/espal-module-core/routes/api/v1/account/overview"
	catalogRoute "github.com/espal-digital-development/espal-module-core/routes/catalog"
	rootRoute "github.com/espal-digital-development/espal-module-core/routes/root"
	spaRoute "github.com/espal-digital-development/espal-module-core/routes/spa"
	"github.com/juju/errors"
)

var errResolveModulePath = errors.New("failed to resolve module path")

// TODO :: 777777
// - How to hook into existing functionality like Slugs and other Database/Repository functionality
//   - How to get functionality báck into the modules? Some kind of reverse registration injection with interface{}'s?
// - CompatibilityDefintion should describe what versions of the core app works with and
//   whether it colides with other functionality being present in the system through other modules
// - Modules should be able to talk to each other and affect load order

// New returns a new instance of Module.
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
			PublicRootFilesPath: filepath.FromSlash(modulePath + "/app/assets/files/root"),
			ImagesPath:          filepath.FromSlash(modulePath + "/app/assets/images"),
			StylesheetsPath:     filepath.FromSlash(modulePath + "/app/assets/css"),
			JavaScriptPath:      filepath.FromSlash(modulePath + "/app/assets/js"),
		})
		if err != nil {
			return errors.Trace(err)
		}
		m.SetAssets(assets)
		return nil
	}

	config.PreGetTranslationsCallback = func(m modules.Modular) error {
		translations, err := translations.New(&translations.Config{
			Path: filepath.FromSlash(modulePath + "/app/translations"),
		})
		if err != nil {
			return errors.Trace(err)
		}
		m.SetTranslations(translations)
		return nil
	}

	config.PreGetRoutesCallback = func(m modules.Modular) error {
		routes, err := routes.New(&routes.Config{
			Entries: map[string]routes.Handler{
				"/": rootRoute.New(root.New()),

				"/Spa": spaRoute.New(spa.New()),

				"/API/V1/Account":                         apiAccountOverviewRoute.New(m.GetStores().User()),
				"/API/V1/Login":                           apiAccountLoginRoute.New(m.GetStores().User()),
				"/API/V1/Account/Register":                &apiEndPointNotImplemented{},
				"/API/V1/Account/ForgotPassword":          &apiEndPointNotImplemented{},
				"/API/V1/Account/ForgotPasswordSucceeded": &apiEndPointNotImplemented{},
				"/API/V1/Account/PasswordRecovery":        &apiEndPointNotImplemented{},

				"/Catalog": catalogRoute.New(catalog.New()),
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
	if _, err := context.WriteString("This endpoint is not implemented yet."); err != nil {
		context.SetStatusCode(http.StatusInternalServerError)
	}
}
