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
	catalogPage "github.com/espal-digital-development/espal-module-core/pages/catalog"
	rootPage "github.com/espal-digital-development/espal-module-core/pages/root"
	spaPage "github.com/espal-digital-development/espal-module-core/pages/spa"
	"github.com/espal-digital-development/espal-module-core/routes/catalog"
	"github.com/espal-digital-development/espal-module-core/routes/root"
	"github.com/espal-digital-development/espal-module-core/routes/spa"
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

	meta, err := meta.New(&meta.Config{
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

	assets, err := assets.New(&assets.Config{
		PublicRootFilesPath: filepath.FromSlash(modulePath + "/app/assets/files/root"),
		ImagesPath:          filepath.FromSlash(modulePath + "/app/assets/images"),
		StylesheetsPath:     filepath.FromSlash(modulePath + "/app/assets/css"),
		JavaScriptPath:      filepath.FromSlash(modulePath + "/app/assets/js"),
	})
	if err != nil {
		return nil, errors.Trace(err)
	}

	translations, err := translations.New(&translations.Config{
		Path: filepath.FromSlash(modulePath + "/app/translations"),
	})
	if err != nil {
		return nil, errors.Trace(err)
	}

	routes, err := routes.New(&routes.Config{
		Entries: map[string]routes.Handler{
			"/": root.New(rootPage.New()),

			"/Spa": spa.New(spaPage.New()),

			// "/API/V1/Account":                         overview.New(),
			// "/API/V1/Login":                           login.New(),
			"/API/V1/Account/Register":                &apiEndPointNotImplemented{},
			"/API/V1/Account/ForgotPassword":          &apiEndPointNotImplemented{},
			"/API/V1/Account/ForgotPasswordSucceeded": &apiEndPointNotImplemented{},
			"/API/V1/Account/PasswordRecovery":        &apiEndPointNotImplemented{},

			"/Catalog": catalog.New(catalogPage.New()),
		},
	})
	if err != nil {
		return nil, errors.Trace(err)
	}

	return modules.New(&modules.Config{
		MetaDefinition:       meta,
		AssetsProvider:       assets,
		TranslationsProvider: translations,
		RoutesProvider:       routes,
	})
}

type apiEndPointNotImplemented struct{}

func (a *apiEndPointNotImplemented) Handle(context contexts.Context) {
	if _, err := context.WriteString("This endpoint is not implemented yet."); err != nil {
		context.SetStatusCode(http.StatusInternalServerError)
	}
}
