package forgot

import (
	"github.com/espal-digital-development/espal-core/validators"
	"github.com/espal-digital-development/espal-core/validators/formview"
	"github.com/juju/errors"
)

var _ Factory = &Forms{}

type language interface {
	ID() uint16
	Code() string
}

// Factory represents an object that serves new forms.
type Factory interface {
	New(language language) (Form, error)
}

// Forms holds all logic to spawn forms for this type.
type Forms struct {
	validatorsFactory validators.Factory
}

// New creates a new Form instance with the required logic.
func (f *Forms) New(language language) (Form, error) {
	validator, err := f.validatorsFactory.NewForm(language)
	if err != nil {
		return nil, errors.Trace(err)
	}

	emailField := validator.NewEmailField("email")
	emailField.SetValidate()
	emailField.SetHideLabel()
	emailField.SetNeedsToBeEqualToField("repeatEmail")
	if err := validator.AddField(emailField); err != nil {
		return nil, errors.Trace(err)
	}

	repeatEmailField := validator.NewEmailField("repeatEmail")
	repeatEmailField.SetHideLabel()
	if err := validator.AddField(repeatEmailField); err != nil {
		return nil, errors.Trace(err)
	}

	return &Forgot{
		validator: validator,
		view:      formview.New(validator),
	}, nil
}

// New returns a new instance of LoginForm.
func New(validatorsFactory validators.Factory) *Forms {
	return &Forms{
		validatorsFactory: validatorsFactory,
	}
}
