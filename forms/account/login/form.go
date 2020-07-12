package login

import (
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/user"
	"github.com/espal-digital-development/espal-core/validators"
	"github.com/espal-digital-development/espal-core/validators/formview"
	"github.com/juju/errors"
	"golang.org/x/crypto/bcrypt"
)

var _ Form = &Login{}

type translator interface {
	Translate(string) string
}

type routeCtx interface {
	contexts.FormContext
	contexts.AuthenticationContext
	contexts.RequestContext
	translator
}

// Form represents an object that offers typical web form interaction.
type Form interface {
	Submit(context routeCtx) (isSubmitted bool, isValid bool, err error)
	GetUserID() string
	RememberMe() bool
	View() formview.View
	Close()
}

// Login web form.
type Login struct {
	validator validators.Validator
	userStore user.Store
	view      formview.View
	isClosed  bool
	userID    string
}

// Submit will submit and validate the form and handle all the rules.
func (l *Login) Submit(context routeCtx) (isSubmitted bool, isValid bool, err error) {
	if l.isClosed {
		err = errors.Errorf("form is already closed")
		return
	}
	if err = l.validator.HandleFromRequest(context); err != nil {
		return
	}
	isSubmitted = l.validator.IsSubmitted()
	if !isSubmitted {
		return
	}
	if isSubmitted {
		isValid, err = l.validator.IsValid()
		if err != nil {
			return
		}
	}
	if !isValid {
		return
	}
	if isValid, err = l.process(context); err != nil {
		return
	}
	return
}

func (l *Login) process(context translator) (bool, error) {
	user, ok, err := l.userStore.GetOneIDAndPasswordForActiveByEmail(l.validator.Field("email").Value())
	if err != nil {
		return false, errors.Trace(err)
	}
	if !ok {
		l.validator.AddError(context.Translate("theSuppliedCredentialsAreNotValid"))
		return false, nil
	}
	isValid, err := l.validator.IsValid()
	if err != nil || !isValid {
		return isValid, errors.Trace(err)
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password()),
		[]byte(l.validator.Field("password").Value())); err != nil {
		l.validator.AddError(context.Translate("theSuppliedCredentialsAreNotValid"))
		return false, nil
	}
	isValid, err = l.validator.IsValid()
	if err != nil || !isValid {
		return isValid, errors.Trace(err)
	}
	l.userID = user.ID()
	return true, nil
}

// GetUserID returns the User ID that was resolved while submitting the form.
func (l *Login) GetUserID() string {
	return l.userID
}

// RememberMe returns an indicator if the user wants to stay logged in longer.
func (l *Login) RememberMe() bool {
	return l.validator.Field("rememberMe").Value() == "1"
}

// View returns the FormView internal to help render inside html output.
func (l *Login) View() formview.View {
	return l.view
}

// Close will release internals.
func (l *Login) Close() {
	l.validator = nil
	l.userStore = nil
	l.view = nil
}
