package change

import (
	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/user"
	"github.com/espal-digital-development/espal-core/validators"
	"github.com/espal-digital-development/espal-core/validators/formview"
	"github.com/juju/errors"
	"golang.org/x/crypto/bcrypt"
)

var _ Form = &Change{}

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
	View() formview.View
	FormFieldValue(name string) string
	Close()
}

// Change web form.
type Change struct {
	validator validators.Validator
	user      user.UserEntity
	view      formview.View
	isClosed  bool
}

// Submit will submit and validate the form and handle all the rules.
func (c *Change) Submit(context routeCtx) (isSubmitted bool, isValid bool, err error) {
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
	if isValid, err = c.process(context); err != nil {
		return
	}
	return
}

func (c *Change) process(translator translator) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(c.user.Password()),
		[]byte(c.validator.Field("currentPassword").Value())); err != nil {
		c.validator.AddError(translator.Translate("yourCurrentPasswordDidNotMatch"))
	}
	isValid, err := c.validator.IsValid()
	if err != nil || !isValid {
		return isValid, errors.Trace(err)
	}
	return true, nil
}

// View returns the FormView internal to help render inside html output.
func (c *Change) View() formview.View {
	return c.view
}

// FormFieldValue returns a single form value that was resolved while submitting the form.
func (c *Change) FormFieldValue(name string) string {
	return c.validator.FieldValue(name)
}

// Close will release internals.
func (c *Change) Close() {
	c.validator = nil
	c.view = nil
}
