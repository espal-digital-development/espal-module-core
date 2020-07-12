package recovery

import (
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/user"
	"github.com/espal-digital-development/espal-core/validators"
	"github.com/espal-digital-development/espal-core/validators/formview"
	"github.com/juju/errors"
	"golang.org/x/crypto/bcrypt"
)

var _ Form = &Recovery{}

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
	GetPasswordResetCount() *uint8
	FormFieldValue(name string) string
	View() formview.View
	Close()
}

// Recovery web form.
type Recovery struct {
	validator          validators.Validator
	userStore          user.Store
	view               formview.View
	isClosed           bool
	userID             string
	passwordResetCount *uint8
}

// Submit will submit and validate the form and handle all the rules.
func (r *Recovery) Submit(context routeCtx) (isSubmitted bool, isValid bool, err error) {
	if r.isClosed {
		err = errors.Errorf("form is already closed")
		return
	}
	if err = r.validator.HandleFromRequest(context); err != nil {
		return
	}
	isSubmitted = r.validator.IsSubmitted()
	if !isSubmitted {
		return
	}
	if isSubmitted {
		isValid, err = r.validator.IsValid()
		if err != nil {
			return
		}
	}
	if !isValid {
		return
	}
	if isValid, err = r.process(context); err != nil {
		return
	}
	return
}

func (r *Recovery) process(translator translator) (bool, error) {
	user, ok, err := r.userStore.GetOneByEmail(r.validator.Field("email").Value())
	if err != nil {
		return false, errors.Trace(err)
	}
	if !ok {
		r.validator.AddError(translator.Translate("theSuppliedInformationDoesNotMatchAnyActiveAccount"))
		return false, nil
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password()),
		[]byte(r.validator.Field("newPassword").Value())); err == nil {
		r.validator.AddError(translator.Translate("yourPasswordShouldNotBeTheSameAsBefore"))
		return false, nil
	}
	r.userID = user.ID()
	r.passwordResetCount = user.PasswordResetCount()
	return true, nil
}

// GetUserID returns the User ID that was resolved while submitting the form.
func (r *Recovery) GetUserID() string {
	return r.userID
}

// GetPasswordResetCount returns the User PasswordResetCount that was resolved while submitting the form.
func (r *Recovery) GetPasswordResetCount() *uint8 {
	return r.passwordResetCount
}

// FormFieldValue returns a single form value that was resolved while submitting the form.
func (r *Recovery) FormFieldValue(name string) string {
	return r.validator.FieldValue(name)
}

// View returns the FormView internal to help render inside html output.
func (r *Recovery) View() formview.View {
	return r.view
}

// Close will release internals.
func (r *Recovery) Close() {
	r.validator = nil
	r.userStore = nil
	r.view = nil
}
