package core

import (
	"path"
	"path/filepath"
	"runtime"

	spaPage "github.com/espal-digital-development/espal-core/app/modules/core/app/pages/spa"
	"github.com/espal-digital-development/espal-core/app/modules/core/app/routes/spa"
	"github.com/espal-digital-development/espal-core/modules"
	"github.com/espal-digital-development/espal-core/modules/assets"
	"github.com/espal-digital-development/espal-core/modules/meta"
	"github.com/espal-digital-development/espal-core/modules/routes"
	"github.com/espal-digital-development/espal-core/modules/translations"
	"github.com/juju/errors"
)

var errResolveModulePath = errors.New("failed to resolve module path")

// TODO :: 777777
// - How to hook into existing functionality like Slugs and other Database/Repository functionality
//   - How to get functionality báck into the modules? Some kind of reverse registration injection with interface{}'s?
// - CompatibilityDefintion should describe what versions of the core app works with and
//   whether it colides with other functionality being present in the system through other modules

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
			"/Spa": spa.New(spaPage.New()),
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
