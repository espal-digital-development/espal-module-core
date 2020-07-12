package createupdate

import (
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/validators"
	"github.com/espal-digital-development/espal-core/validators/formview"
	"github.com/juju/errors"
)

var _ Form = &CreateUpdate{}

type context interface {
	contexts.RequestContext
	contexts.AuthenticationContext
	contexts.FormContext
}

// Form represents an object that offers typical web form interaction.
type Form interface {
	Submit(context context) (isSubmitted bool, isValid bool, err error)
	View() formview.View
	FieldValue(name string) string
	FieldPointerValue(name string) *string
	FieldValueAsUint(name string) uint
	Close()
}

// CreateUpdate web form.
type CreateUpdate struct {
	validator validators.Validator
	view      formview.View
	isClosed  bool
}

// Submit will submit and validate the form and handle all the rules.
func (c *CreateUpdate) Submit(context context) (isSubmitted bool, isValid bool, err error) {
	if c.isClosed {
		err = errors.Errorf("form is already closed")
		return
	}
	if err = c.validator.HandleFromRequest(context); err != nil {
		return
	}
	isSubmitted = c.validator.IsSubmitted()
	if !isSubmitted {
		return
	}
	if isSubmitted {
		isValid, err = c.validator.IsValid()
		if err != nil {
			return
		}
	}
	return
}

// View returns the FormView internal to help render inside html output.
func (c *CreateUpdate) View() formview.View {
	return c.view
}

// FieldValue returns a single form value that was resolved while submitting the form.
func (c *CreateUpdate) FieldValue(name string) string {
	return c.validator.FieldValue(name)
}

// FieldPointerValue returns a single form pointer value that was resolved while submitting the form.
func (c *CreateUpdate) FieldPointerValue(name string) *string {
	return c.validator.Field(name).PointerValue()
}

// FieldValueAsUint returns a single form uint value that was resolved while submitting the form.
func (c *CreateUpdate) FieldValueAsUint(name string) uint {
	return c.validator.Field(name).ValueAsUint()
}

// Close will release internals.
func (c *CreateUpdate) Close() {
	c.validator = nil
	c.view = nil
}
