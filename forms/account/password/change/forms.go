package change

import (
	"github.com/espal-digital-development/espal-core/stores/user"
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
	New(language language, user user.UserEntity) (Form, error)
}

// Forms holds all logic to spawn forms for this type.
type Forms struct {
	validatorsFactory validators.Factory
}

// New creates a new Form instance with the required logic.
func (f *Forms) New(language language, user user.UserEntity) (Form, error) {
	validator, err := f.validatorsFactory.NewForm(language)
	if err != nil {
		return nil, errors.Trace(err)
	}

	if err := validator.AddField(validator.NewPasswordField("currentPassword")); err != nil {
		return nil, errors.Trace(err)
	}

	newPasswordField := validator.NewPasswordField("newPassword")
	newPasswordField.SetValidate()
	newPasswordField.SetCannotBeEqualToField("currentPassword")
	newPasswordField.SetNeedsToBeEqualToField("repeatNewPassword")
	if err := validator.AddField(newPasswordField); err != nil {
		return nil, errors.Trace(err)
	}

	if err := validator.AddField(validator.NewPasswordField("repeatNewPassword")); err != nil {
		return nil, errors.Trace(err)
	}

	return &Change{
		validator: validator,
		user:      user,
		view:      formview.New(validator),
	}, nil
}

// New returns a new instance of LoginForm.
func New(validatorsFactory validators.Factory) *Forms {
	return &Forms{
		validatorsFactory: validatorsFactory,
	}
}
