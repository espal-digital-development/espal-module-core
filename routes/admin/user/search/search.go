package search

import (
	"fmt"

	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/user"
	"github.com/juju/errors"
)

// Route processor.
type Route struct {
	userStore user.Store
}

// Handle route handler.
func (r *Route) Handle(context contexts.Context) {
	if !context.HasUserRightOrForbid("ReadUser") {
		return
	}

	users, _, err := r.userStore.Filter(context)
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}

	var firsthad bool
	for k := range users {
		// TODO :: This output format should get more variations base on settings
		if firsthad {
			if _, err := context.WriteString("\n"); err != nil {
				context.RenderInternalServerError(errors.Trace(err))
				return
			}
		}
		if _, err := context.WriteString(fmt.Sprintf("%s\t", users[k].ID())); err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
		if users[k].FirstName() != nil {
			if _, err := context.WriteString(*users[k].FirstName()); err != nil {
				context.RenderInternalServerError(errors.Trace(err))
				return
			}
		}
		if users[k].Surname() != nil {
			if users[k].Surname() != nil {
				if _, err := context.WriteString(" "); err != nil {
					context.RenderInternalServerError(errors.Trace(err))
					return
				}
			}
			if _, err := context.WriteString(*users[k].Surname()); err != nil {
				context.RenderInternalServerError(errors.Trace(err))
				return
			}
		}
		if users[k].FirstName() == nil && users[k].Surname() == nil {
			if _, err := context.WriteString(users[k].Email()); err != nil {
				context.RenderInternalServerError(errors.Trace(err))
				return
			}
		} else {
			if _, err := context.WriteString(fmt.Sprintf(" (%s)", users[k].Email())); err != nil {
				context.RenderInternalServerError(errors.Trace(err))
				return
			}
		}
		if !firsthad {
			firsthad = true
		}
	}
}

// New returns a new instance of Route.
func New(userStore user.Store) *Route {
	return &Route{
		userStore: userStore,
	}
}
