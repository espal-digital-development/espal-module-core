package createupdate

import (
	"strconv"
	"strings"
	"time"

	"github.com/espal-digital-development/espal-core/repositories/translations"
	"github.com/espal-digital-development/espal-core/storage"
	"github.com/espal-digital-development/espal-core/stores/user"
	useraddressstore "github.com/espal-digital-development/espal-core/stores/user/address"
	usercontactstore "github.com/espal-digital-development/espal-core/stores/user/contact"
	"github.com/espal-digital-development/espal-core/validators"
	"github.com/espal-digital-development/espal-core/validators/formview"
	"github.com/juju/errors"
)

var _ Factory = &Forms{}

const maxHumanAgeRange = 120

type language interface {
	ID() uint16
	Code() string
}

// Factory represents an object that serves new forms.
type Factory interface {
	New(user user.UserEntity, language language) (Form, error)
}

// Forms holds all logic to spawn forms for this type.
type Forms struct {
	validatorsFactory      validators.Factory
	translationsRepository translations.Repository
	storage                storage.Modifyable
	userAddressStore       useraddressstore.Store
	userContactStore       usercontactstore.Store
}

// New creates a new Form instance with the required logic.
// nolint:gocyclo,funlen
func (f *Forms) New(user user.UserEntity, language language) (Form, error) {
	validator, err := f.validatorsFactory.NewForm(language)
	if err != nil {
		return nil, errors.Trace(err)
	}

	activeField := validator.NewCheckboxField("active")
	if user.ID() != "" && user.Active() {
		activeField.SetValue("1")
	}
	if err := validator.AddField(activeField); err != nil {
		return nil, errors.Trace(err)
	}

	priorityField := validator.NewNumberField("priority")
	priorityField.SetOptional()
	if user.ID() != "" && user.Priority() > 0 {
		priorityField.SetValue(strconv.FormatUint(uint64(user.Priority()), 10))
	}
	if err := validator.AddField(priorityField); err != nil {
		return nil, errors.Trace(err)
	}

	emailField := validator.NewEmailField("email")
	emailField.SetValidate()
	if user.ID() != "" {
		emailField.SetValue(user.Email())
	}
	if err := validator.AddField(emailField); err != nil {
		return nil, errors.Trace(err)
	}

	avatarField := validator.NewFileField("avatar", f.storage)
	avatarField.SetOptional()
	avatarField.SetMaxLength(75)
	avatarField.SetAllowedMIMETypes([]string{"image/jpeg", "image/png"})
	if user.ID() != "" && user.Avatar() != nil {
		avatarField.SetValue(*user.Avatar())
	}
	if err := validator.AddField(avatarField); err != nil {
		return nil, errors.Trace(err)
	}

	firstNameField := validator.NewTextField("firstName")
	avatarField.SetMinLength(1)
	avatarField.SetMaxLength(50)
	avatarField.SetOptional()
	if user.ID() != "" && user.FirstName() != nil {
		firstNameField.SetValue(*user.FirstName())
	}
	if err := validator.AddField(firstNameField); err != nil {
		return nil, errors.Trace(err)
	}

	surnameField := validator.NewTextField("surname")
	surnameField.SetMinLength(1)
	surnameField.SetMaxLength(50)
	surnameField.SetOptional()
	if user.ID() != "" && user.Surname() != nil {
		surnameField.SetValue(*user.Surname())
	}
	if err := validator.AddField(surnameField); err != nil {
		return nil, errors.Trace(err)
	}

	dateOfBirthField := validator.NewDateTimeField("dateOfBirth")
	dateOfBirthField.SetOptional()
	dateOfBirthField.SetMinYear(uint(time.Now().Year() - maxHumanAgeRange))
	if user.ID() != "" && user.DateOfBirth() != nil {
		dateOfBirthField.SetValue(user.DateOfBirth().Format(time.RFC3339[0:9]))
	}
	if err := validator.AddField(dateOfBirthField); err != nil {
		return nil, errors.Trace(err)
	}

	passwordField := validator.NewPasswordField("password")
	passwordField.SetValidate()
	passwordField.SetNeedsToBeEqualToField("repeatPassword")
	passwordField.SetCannotBeEqualToField("email")

	repeatPasswordField := validator.NewPasswordField("repeatPassword")
	if user.ID() != "" {
		passwordField.SetOptional()
		repeatPasswordField.SetOptional()
	}
	if err := validator.AddField(passwordField); err != nil {
		return nil, errors.Trace(err)
	}
	if err := validator.AddField(repeatPasswordField); err != nil {
		return nil, errors.Trace(err)
	}

	countryField := validator.NewChoiceField("country")
	countryField.SetOptional()
	countryField.SetSearchable()
	countryField.SetCheckValuesInChoices()
	countryField.SetChoices(f.validatorsFactory.GetCountryOptionsForLanguage(language))
	if user.ID() != "" && user.Country() != nil {
		countryField.SetValue(strconv.FormatUint(uint64(*user.Country()), 10))
	}
	if err := validator.AddField(countryField); err != nil {
		return nil, errors.Trace(err)
	}

	languageField := validator.NewChoiceField("language")
	languageField.SetSearchable()
	languageField.SetCheckValuesInChoices()
	languageField.SetChoices(f.validatorsFactory.GetLanguageOptionsForLanguage(language))
	if user.ID() != "" {
		languageField.SetValue(strconv.FormatUint(uint64(user.Language()), 10))
	}
	if err := validator.AddField(languageField); err != nil {
		return nil, errors.Trace(err)
	}

	biographyField := validator.NewTextAreaField("biography")
	biographyField.SetOptional()
	biographyField.SetMaxLength(1000)
	if user.ID() != "" && user.Biography() != nil {
		biographyField.SetValue(*user.Biography())
	}
	if err := validator.AddField(biographyField); err != nil {
		return nil, errors.Trace(err)
	}

	commentsField := validator.NewTextAreaField("comments")
	commentsField.SetOptional()
	commentsField.SetMaxLength(1000)
	commentsField.SetPlaceholder(f.translationsRepository.Plural(language.ID(), "comment") + " (" +
		f.translationsRepository.Singular(language.ID(), "optional") + ")")
	commentsField.SetDontTranslate()
	if user.ID() != "" && user.Comments() != nil {
		commentsField.SetValue(*user.Comments())
	}
	if err := validator.AddField(commentsField); err != nil {
		return nil, errors.Trace(err)
	}

	currenciesField := validator.NewChoiceField("currencies")
	currenciesField.SetOptional()
	currenciesField.SetPlaceholder(f.translationsRepository.Plural(language.ID(), "currency") + " (" +
		f.translationsRepository.Singular(language.ID(), "optional") + ")")
	currenciesField.SetDontTranslate()
	currenciesField.SetSearchable()
	currenciesField.SetCheckValuesInChoices()
	currenciesField.SetChoices(f.validatorsFactory.GetCurrencyOptionsForLanguage(language))
	currenciesField.SetMultiple()
	if user.ID() != "" {
		if err := currenciesField.SetValues(strings.Split(user.Currencies(), ",")); err != nil {
			return nil, errors.Trace(err)
		}
	}
	if err := validator.AddField(currenciesField); err != nil {
		return nil, errors.Trace(err)
	}

	if user.ID() != "" {
		var deliveryAddresses []validators.ChoiceOption
		var invoiceAddresses []validators.ChoiceOption
		addresses, ok, err := f.userAddressStore.ForUser(user.ID())
		if err != nil {
			return nil, errors.Trace(err)
		}
		if ok {
			for k := range addresses {
				option := f.validatorsFactory.NewChoiceOption(addresses[k].ID(),
					f.userAddressStore.DisplayValue(addresses[k], language.ID()))
				deliveryAddresses = append(deliveryAddresses, option)
				invoiceAddresses = append(invoiceAddresses, option)
			}
		}

		deliveryAddressField := validator.NewChoiceField("defaultDeliveryAddress")
		deliveryAddressField.SetOptional()
		deliveryAddressField.SetChoices(deliveryAddresses)
		if user.DefaultDeliveryAddressID() != nil {
			deliveryAddressField.SetValue(*user.DefaultDeliveryAddressID())
		}
		if err := validator.AddField(deliveryAddressField); err != nil {
			return nil, errors.Trace(err)
		}

		invoiceAddressField := validator.NewChoiceField("defaultInvoiceAddress")
		invoiceAddressField.SetOptional()
		invoiceAddressField.SetChoices(invoiceAddresses)
		if user.DefaultInvoiceAddressID() != nil {
			invoiceAddressField.SetValue(*user.DefaultInvoiceAddressID())
		}
		if err := validator.AddField(invoiceAddressField); err != nil {
			return nil, errors.Trace(err)
		}
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
func New(validatorsFactory validators.Factory, translationsRepository translations.Repository,
	storage storage.Modifyable, userAddressStore useraddressstore.Store,
	userContactStore usercontactstore.Store) *Forms {
	return &Forms{
		validatorsFactory:      validatorsFactory,
		translationsRepository: translationsRepository,
		storage:                storage,
		userAddressStore:       userAddressStore,
		userContactStore:       userContactStore,
	}
}
