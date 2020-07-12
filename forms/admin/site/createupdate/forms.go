package createupdate

import (
	"strconv"
	"strings"

	"github.com/espal-digital-development/espal-core/repositories/translations"
	"github.com/espal-digital-development/espal-core/validators"
	"github.com/espal-digital-development/espal-core/validators/formview"
	"github.com/juju/errors"
)

var _ Factory = &Forms{}

type language interface {
	ID() uint16
	Code() string
}

type site interface {
	ID() string
	Online() bool
	Country() *uint16
	Language() *uint16
	Currencies() string
}

// Factory represents an object that serves new forms.
type Factory interface {
	New(site site, language language) (Form, error)
}

// Forms holds all logic to spawn forms for this type.
type Forms struct {
	validatorsFactory      validators.Factory
	translationsRepository translations.Repository
}

// New creates a new Form instance with the required logic.
func (f *Forms) New(site site, language language) (Form, error) {
	validator, err := f.validatorsFactory.NewForm(language)
	if err != nil {
		return nil, errors.Trace(err)
	}

	onlineField := validator.NewCheckboxField("online")
	if site.ID() != "" && site.Online() {
		onlineField.SetValue("1")
	}
	if err := validator.AddField(onlineField); err != nil {
		return nil, errors.Trace(err)
	}

	countryField := validator.NewChoiceField("country")
	countryField.SetOptional()
	countryField.SetSearchable()
	countryField.SetCheckValuesInChoices()
	countryField.SetChoices(f.validatorsFactory.GetCountryOptionsForLanguage(language))
	if site.ID() != "" && site.Country() != nil {
		countryField.SetValue(strconv.FormatUint(uint64(*site.Country()), 10))
	}
	if err := validator.AddField(countryField); err != nil {
		return nil, errors.Trace(err)
	}

	languageField := validator.NewChoiceField("language")
	languageField.SetSearchable()
	languageField.SetCheckValuesInChoices()
	languageField.SetChoices(f.validatorsFactory.GetLanguageOptionsForLanguage(language))
	if site.ID() != "" && site.Language() != nil {
		languageField.SetValue(strconv.FormatUint(uint64(*site.Language()), 10))
	}
	if err := validator.AddField(languageField); err != nil {
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
	if site.ID() != "" {
		if err := currenciesField.SetValues(strings.Split(site.Currencies(), ",")); err != nil {
			return nil, errors.Trace(err)
		}
	}
	if err := validator.AddField(currenciesField); err != nil {
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
func New(validatorsFactory validators.Factory, translationsRepository translations.Repository) *Forms {
	return &Forms{
		validatorsFactory:      validatorsFactory,
		translationsRepository: translationsRepository,
	}
}
