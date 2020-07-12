package view

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/espal-digital-development/espal-core/pageactions"
	"github.com/espal-digital-development/espal-core/repositories/userrights"
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/user/group"
	page "github.com/espal-digital-development/espal-module-core/pages/admin/user/group/view"
	"github.com/juju/errors"
)

// Route processor.
type Route struct {
	userRightsRepository userrights.Repository
	userGroupStore       group.Store
	viewPageFactory      page.Factory
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	if !context.HasUserRightOrForbid("ReadUserGroup") {
		return
	}

	id := context.QueryValue("id")
	if id == "" {
		context.RenderNotFound()
		return
	}

	userGroup, ok, err := r.userGroupStore.GetOneByIDWithCreator(id)
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	if !ok {
		context.RenderNotFound()
		return
	}

	translations, translationsFound, err := r.userGroupStore.TranslationsForID(userGroup.ID())
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}

	var userRightsList []string
	var userGroupUserRights map[uint16]bool

	if userRightsCount := userGroup.UserRightsCount(); userRightsCount > 0 {
		userRightsList = strings.Split(userGroup.UserRights(), ",")
		userGroupUserRights = make(map[uint16]bool, userRightsCount)
		for k := range userRightsList {
			userRightUint16, err := strconv.ParseUint(userRightsList[k], 10, 16)
			if err != nil {
				context.RenderInternalServerError(errors.Trace(err))
				return
			}
			userGroupUserRights[uint16(userRightUint16)] = true
		}
	}

	var canUpdate bool
	var canDelete bool
	translationsActions := pageactions.New(context, "UserGroup", translationsFound)
	translationsActions.AddCreateWithFieldAndPath("translation", fmt.Sprintf("UserGroup/Translation/Create?id=%s",
		userGroup.ID()))

	if translationsFound {
		translationsActions.AddDeleteWithPath(fmt.Sprintf("UserGroup/Translation/Delete?id=%s", userGroup.ID()))
		canUpdate = context.HasUserRight("UpdateUserGroup")
		canDelete = context.HasUserRight("DeleteUserGroup")
	}

	language, err := context.GetLanguage()
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	pageActions := pageactions.New(context, "UserGroup", len(userRightsList) > 0)
	pageActions.AddUpdateWithFieldAndPath("userright", fmt.Sprintf("UserGroup/UserRights/Update?id=%s", userGroup.ID()))
	r.viewPageFactory.NewPage(
		context,
		language,
		userGroup,
		r.userRightsRepository.AllByCode(),
		r.userRightsRepository.UserRightCodes(),
		userGroupUserRights,
		pageActions,
		translations,
		translationsActions,
		canUpdate,
		canDelete).Render()
}

// New returns a new instance of Route.
func New(userRightsRepository userrights.Repository, userGroupStore group.Store, viewPageFactory page.Factory) *Route {
	return &Route{
		userRightsRepository: userRightsRepository,
		userGroupStore:       userGroupStore,
		viewPageFactory:      viewPageFactory,
	}
}
