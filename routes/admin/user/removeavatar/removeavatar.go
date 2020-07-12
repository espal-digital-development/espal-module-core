package removeavatar

import (
	"net/http"

	"github.com/espal-digital-development/espal-core/config"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/storage"
	"github.com/espal-digital-development/espal-core/stores/user"
	"github.com/juju/errors"
)

// Route processor.
type Route struct {
	configService            config.Config
	assetsPublicFilesStorage storage.Modifyable
	userStore                user.Store
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	if !context.HasUserRightOrForbid("UpdateUser") {
		return
	}

	id := context.QueryValue("id")
	if len(id) == 0 {
		context.RenderNotFound()
		return
	}

	avatar, ok, err := r.userStore.GetAvatar(id)
	if err != nil {
		if err := context.SetFlashErrorMessage(err.Error()); err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
	}
	if !ok {
		context.RenderNotFound()
		return
	}
	if avatar == nil {
		if err := context.SetFlashWarningMessage(context.Translate("avatarIsNotSet")); err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
	} else {
		if r.assetsPublicFilesStorage.Exists(*avatar + ".gz") {
			if err := r.assetsPublicFilesStorage.Delete(*avatar + ".gz"); err != nil {
				context.RenderInternalServerError(errors.Trace(err))
				return
			}
		}
		if r.assetsPublicFilesStorage.Exists(*avatar) {
			if err := r.assetsPublicFilesStorage.Delete(*avatar); err != nil {
				context.RenderInternalServerError(errors.Trace(err))
				return
			}
		}

		if err = r.userStore.UnsetAvatar(id); err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
	}

	if err := context.SetFlashSuccessMessage(context.Translate("removalWasSuccessful")); err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}

	if referer := context.Referer(); len(referer) > 0 {
		context.Redirect(referer, http.StatusTemporaryRedirect)
	} else {
		context.Redirect(context.AdminURL()+"", http.StatusTemporaryRedirect)
	}
}

// New returns a new instance of Route.
func New(configService config.Config, assetsPublicFilesStorage storage.Modifyable, userStore user.Store) *Route {
	return &Route{
		configService:            configService,
		assetsPublicFilesStorage: assetsPublicFilesStorage,
		userStore:                userStore,
	}
}
