package forgot

import (
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/validators"
	"github.com/espal-digital-development/espal-core/validators/formview"
	"github.com/juju/errors"
)

var _ Form = &Forgot{}

type context interface {
	contexts.RequestContext
	contexts.AuthenticationContext
	contexts.FormContext
}

// Form represents an object that offers typical web form interaction.
type Form interface {
	Submit(context context) (isSubmitted bool, isValid bool, err error)
	FormFieldValue(name string) string
	View() formview.View
	Close()
}

// Forgot web form.
type Forgot struct {
	validator validators.Validator
	view      formview.View
	isClosed  bool
}

// Submit will submit and validate the form and handle all the rules.
func (f *Forgot) Submit(context context) (isSubmitted bool, isValid bool, err error) {
	if f.isClosed {
		err = errors.Errorf("form is already closed")
		return
	}
	if err = f.validator.HandleFromRequest(context); err != nil {
		return
	}
	isSubmitted = f.validator.IsSubmitted()
	if !isSubmitted {
		return
	}
	if isSubmitted {
		isValid, err = f.validator.IsValid()
		if err != nil {
			return
		}
	}
	return
}

// FormFieldValue returns a single form value that was resolved while submitting the form.
func (f *Forgot) FormFieldValue(name string) string {
	return f.validator.FieldValue(name)
}

// View returns the FormView internal to help render inside html output.
func (f *Forgot) View() formview.View {
	return f.view
}

// Close will release internals.
func (f *Forgot) Close() {
	f.validator = nil
	f.view = nil
}
