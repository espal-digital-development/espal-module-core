package register

import (
	"database/sql"

	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/user"
	"github.com/espal-digital-development/espal-core/validators"
	"github.com/espal-digital-development/espal-core/validators/formview"
	"github.com/juju/errors"
)

var _ Form = &Register{}

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
	FormFieldValue(name string) string
	View() formview.View
	Close()
}

// Register web form.
type Register struct {
	validator validators.Validator
	userStore user.Store
	view      formview.View
	isClosed  bool
}

// Submit will submit and validate the form and handle all the rules.
func (r *Register) Submit(context routeCtx) (isSubmitted bool, isValid bool, err error) {
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

func (r *Register) process(translator translator) (bool, error) {
	exists, err := r.userStore.ExistsByEmail(r.validator.Field("email").Value())
	if err != nil && err != sql.ErrNoRows {
		return false, errors.Trace(err)
	}
	if exists {
		r.validator.Field("email").AddError(translator.Translate("emailIsAlreadyUsed"))
	}
	return r.validator.IsValid()
}

// FormFieldValue returns a single form value that was resolved while submitting the form.
func (r *Register) FormFieldValue(name string) string {
	return r.validator.FieldValue(name)
}

// View returns the FormView internal to help render inside html output.
func (r *Register) View() formview.View {
	return r.view
}

// Close will release internals.
func (r *Register) Close() {
	r.validator = nil
	r.userStore = nil
	r.view = nil
}
