package createupdate

import (
	"strconv"

	"github.com/espal-digital-development/espal-core/validators"
	"github.com/espal-digital-development/espal-core/validators/formview"
	"github.com/juju/errors"
)

var _ Factory = &Forms{}

type language interface {
	ID() uint16
	Code() string
}

type address interface {
	ID() string
	Active() bool
	FirstName() *string
	Surname() *string
	Street() string
	StreetLine2() *string
	Number() string
	NumberAddition() *string
	ZipCode() string
	City() string
	State() *string
	Country() *uint16
	PhoneNumber() *string
	Email() *string
}

// Factory represents an object that serves new forms.
type Factory interface {
	New(address address, language language) (Form, error)
}

// Forms holds all logic to spawn forms for this type.
type Forms struct {
	validatorsFactory validators.Factory
}

// New creates a new Form instance with the required logic.
// nolint:gocyclo,funlen
func (f *Forms) New(address address, language language) (Form, error) {
	validator, err := f.validatorsFactory.NewForm(language)
	if err != nil {
		return nil, errors.Trace(err)
	}

	activeField := validator.NewCheckboxField("active")
	if (address.ID() != "" && address.Active()) || address.ID() == "" {
		activeField.SetValue("1")
	}
	if err := validator.AddField(activeField); err != nil {
		return nil, errors.Trace(err)
	}

	firstNameField := validator.NewTextField("firstName")
	firstNameField.SetOptional()
	firstNameField.SetMinLength(1)
	firstNameField.SetMaxLength(50)
	if address.ID() != "" && address.FirstName() != nil {
		firstNameField.SetValue(*address.FirstName())
	}
	if err := validator.AddField(firstNameField); err != nil {
		return nil, errors.Trace(err)
	}

	surnameField := validator.NewTextField("surname")
	surnameField.SetOptional()
	surnameField.SetMinLength(1)
	surnameField.SetMaxLength(50)
	if address.ID() != "" && address.Surname() != nil {
		surnameField.SetValue(*address.Surname())
	}
	if err := validator.AddField(surnameField); err != nil {
		return nil, errors.Trace(err)
	}

	streetField := validator.NewTextField("street")
	streetField.SetMinLength(1)
	streetField.SetMaxLength(70)
	if address.ID() != "" {
		streetField.SetValue(address.Street())
	}
	if err := validator.AddField(streetField); err != nil {
		return nil, errors.Trace(err)
	}

	streetLine2Field := validator.NewTextField("streetLine2")
	streetLine2Field.SetOptional()
	streetLine2Field.SetMinLength(1)
	streetLine2Field.SetMaxLength(50)
	if address.ID() != "" && address.StreetLine2() != nil {
		streetLine2Field.SetValue(*address.StreetLine2())
	}
	if err := validator.AddField(streetLine2Field); err != nil {
		return nil, errors.Trace(err)
	}

	numberField := validator.NewTextField("number")
	numberField.SetMinLength(1)
	numberField.SetMaxLength(50)
	if address.ID() != "" {
		numberField.SetValue(address.Number())
	}
	if err := validator.AddField(numberField); err != nil {
		return nil, errors.Trace(err)
	}

	numberAdditionField := validator.NewTextField("numberAddition")
	numberAdditionField.SetOptional()
	numberAdditionField.SetMinLength(1)
	numberAdditionField.SetMaxLength(50)
	if address.ID() != "" && address.NumberAddition() != nil {
		numberAdditionField.SetValue(*address.NumberAddition())
	}
	if err := validator.AddField(numberAdditionField); err != nil {
		return nil, errors.Trace(err)
	}

	zipCodeField := validator.NewTextField("zipCode")
	zipCodeField.SetMinLength(1)
	zipCodeField.SetMaxLength(50)
	if address.ID() != "" {
		zipCodeField.SetValue(address.ZipCode())
	}
	if err := validator.AddField(zipCodeField); err != nil {
		return nil, errors.Trace(err)
	}

	cityField := validator.NewTextField("city")
	cityField.SetMinLength(1)
	cityField.SetMaxLength(50)
	if address.ID() != "" {
		cityField.SetValue(address.City())
	}
	if err := validator.AddField(cityField); err != nil {
		return nil, errors.Trace(err)
	}

	stateField := validator.NewTextField("state")
	stateField.SetOptional()
	stateField.SetMinLength(1)
	stateField.SetMaxLength(50)
	if address.ID() != "" && address.State() != nil {
		stateField.SetValue(*address.State())
	}
	if err := validator.AddField(stateField); err != nil {
		return nil, errors.Trace(err)
	}

	countryField := validator.NewChoiceField("country")
	countryField.SetOptional()
	countryField.SetSearchable()
	countryField.SetCheckValuesInChoices()
	countryField.SetChoices(f.validatorsFactory.GetCountryOptionsForLanguage(language))
	if address.ID() != "" && address.Country() != nil {
		countryField.SetValue(strconv.FormatUint(uint64(*address.Country()), 10))
	}
	if err := validator.AddField(countryField); err != nil {
		return nil, errors.Trace(err)
	}

	phoneNumberField := validator.NewTextField("phoneNumber")
	phoneNumberField.SetOptional()
	phoneNumberField.SetMinLength(1)
	phoneNumberField.SetMaxLength(50)
	if address.ID() != "" && address.PhoneNumber() != nil {
		phoneNumberField.SetValue(*address.PhoneNumber())
	}
	if err := validator.AddField(phoneNumberField); err != nil {
		return nil, errors.Trace(err)
	}

	emailField := validator.NewEmailField("email")
	emailField.SetOptional()
	emailField.SetValidate()
	emailField.SetMinLength(7)
	emailField.SetMaxLength(255)
	if address.ID() != "" && address.Email() != nil {
		emailField.SetValue(*address.Email())
	}
	if err := validator.AddField(emailField); err != nil {
		return nil, errors.Trace(err)
	}

	actionField := validator.NewHiddenField("action")
	actionField.SetOptional()
	if err := validator.AddField(actionField); err != nil {
		return nil, errors.Trace(err)
	}

	return &CreateUpdate{
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
