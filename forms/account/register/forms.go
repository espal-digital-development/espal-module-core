package register

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
	repeatEmailField.SetValidate()
	repeatEmailField.SetHideLabel()
	if err := validator.AddField(repeatEmailField); err != nil {
		return nil, errors.Trace(err)
	}

	passwordField := validator.NewPasswordField("password")
	passwordField.SetValidate()
	passwordField.SetHideLabel()
	passwordField.SetNeedsToBeEqualToField("repeatPassword")
	passwordField.SetCannotBeEqualToField("email")
	if err := validator.AddField(passwordField); err != nil {
		return nil, errors.Trace(err)
	}

	repeatPasswordField := validator.NewPasswordField("repeatPassword")
	repeatPasswordField.SetValidate()
	repeatPasswordField.SetHideLabel()
	if err := validator.AddField(repeatPasswordField); err != nil {
		return nil, errors.Trace(err)
	}

	firstNameField := validator.NewTextField("firstName")
	firstNameField.SetMinLength(1)
	firstNameField.SetMaxLength(50)
	firstNameField.SetOptional()
	firstNameField.SetHideLabel()
	if err := validator.AddField(firstNameField); err != nil {
		return nil, errors.Trace(err)
	}

	surnameField := validator.NewTextField("surname")
	surnameField.SetMinLength(1)
	surnameField.SetMaxLength(50)
	surnameField.SetOptional()
	surnameField.SetHideLabel()
	if err := validator.AddField(surnameField); err != nil {
		return nil, errors.Trace(err)
	}

	dateOfBirthField := validator.NewDateTimeField("dateOfBirth")
	dateOfBirthField.SetMinYear(1900)
	dateOfBirthField.SetOptional()
	dateOfBirthField.SetHideLabel()
	if err := validator.AddField(dateOfBirthField); err != nil {
		return nil, errors.Trace(err)
	}

	return &Register{
		validator: validator,
		userStore: f.userStore,
		view:      formview.New(validator),
	}, nil
}

// New returns a new instance of LoginForm.
func New(validatorsFactory validators.Factory, userStore user.Store) *Forms {
	return &Forms{
		validatorsFactory: validatorsFactory,
		userStore:         userStore,
	}
}
