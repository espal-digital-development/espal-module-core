package login

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
	New(language language) (Form, error)
}

// Forms holds all logic to spawn forms for this type.
type Forms struct {
	validatorsFactory validators.Factory
	userStore         user.Store
}

// New creates a new Form instance with the required logic.
func (f *Forms) New(language language) (Form, error) {
	form, err := f.validatorsFactory.NewForm(language)
	if err != nil {
		return nil, errors.Trace(err)
	}

	emailField := form.NewEmailField("email")
	emailField.SetValidate()
	emailField.SetHideLabel()
	if err := form.AddField(emailField); err != nil {
		return nil, errors.Trace(err)
	}

	passwordField := form.NewPasswordField("password")
	passwordField.SetHideLabel()
	if err := form.AddField(passwordField); err != nil {
		return nil, errors.Trace(err)
	}

	rememberMeField := form.NewCheckboxField("rememberMe")
	rememberMeField.SetHideLabel()
	if err := form.AddField(rememberMeField); err != nil {
		return nil, errors.Trace(err)
	}

	return &Login{
		validator: form,
		userStore: f.userStore,
		view:      formview.New(form),
	}, nil
}

// New returns a new instance of LoginForm.
func New(validatorsFactory validators.Factory, userStore user.Store) *Forms {
	return &Forms{
		validatorsFactory: validatorsFactory,
		userStore:         userStore,
	}
}
