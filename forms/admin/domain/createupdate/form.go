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
	FieldValue(name string) string
	FieldValues(name string) []string
	FieldValueAsBool(name string) bool
	FieldValueAsUint(name string) uint
	FieldValueAsUint16(name string) uint16
	View() formview.View
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
	if !isValid {
		return
	}
	return
}

// FieldValue returns a single form value that was resolved while submitting the form.
func (c *CreateUpdate) FieldValue(name string) string {
	return c.validator.FieldValue(name)
}

// FieldValues returns a single form values that was resolved while submitting the form.
func (c *CreateUpdate) FieldValues(name string) []string {
	return c.validator.Field(name).Values()
}

// FieldValueAsBool returns a single form bool value that was resolved while submitting the form.
func (c *CreateUpdate) FieldValueAsBool(name string) bool {
	return c.validator.Field(name).ValueAsBool()
}

// FieldValueAsUint returns a single form uint value that was resolved while submitting the form.
func (c *CreateUpdate) FieldValueAsUint(name string) uint {
	return c.validator.Field(name).ValueAsUint()
}

// FieldValueAsUint16 returns a single form uint16 value that was resolved while submitting the form.
func (c *CreateUpdate) FieldValueAsUint16(name string) uint16 {
	return c.validator.Field(name).ValueAsUint16()
}

// View returns the FormView internal to help render inside html output.
func (c *CreateUpdate) View() formview.View {
	return c.view
}

// Close will release internals.
func (c *CreateUpdate) Close() {
	c.validator = nil
	c.view = nil
}
