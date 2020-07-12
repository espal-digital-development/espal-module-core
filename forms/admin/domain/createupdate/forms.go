package createupdate

import (
	"strconv"
	"strings"

	"github.com/espal-digital-development/espal-core/config"
	"github.com/espal-digital-development/espal-core/repositories/translations"
	"github.com/espal-digital-development/espal-core/stores/site"
	"github.com/espal-digital-development/espal-core/validators"
	"github.com/espal-digital-development/espal-core/validators/formview"
	"github.com/juju/errors"
)

var _ Factory = &Forms{}

type language interface {
	ID() uint16
	Code() string
}

type domain interface {
	ID() string
	SiteID() string
	Active() bool
	Host() string
	Language() *uint16
	Currencies() string
}

// Factory represents an object that serves new forms.
type Factory interface {
	New(domain domain, language language) (Form, error)
}

// Forms holds all logic to spawn forms for this type.
type Forms struct {
	validatorsFactory      validators.Factory
	translationsRepository translations.Repository
	configService          config.Config
	siteStore              site.Store
}

// New creates a new Form instance with the required logic.
func (f *Forms) New(domain domain, language language) (Form, error) {
	validator, err := f.validatorsFactory.NewForm(language)
	if err != nil {
		return nil, errors.Trace(err)
	}

	activeField := validator.NewCheckboxField("active")
	if domain.ID() != "" && domain.Active() {
		activeField.SetValue("1")
	}
	if err := validator.AddField(activeField); err != nil {
		return nil, errors.Trace(err)
	}

	siteField := validator.NewChoiceField("site")
	siteField.SetSearchable()
	siteField.SetSearchableDataPath(f.configService.AdminURL() + "/Site/Search")
	if err := validator.AddField(siteField); err != nil {
		return nil, errors.Trace(err)
	}

	hostField := validator.NewTextField("host")
	hostField.SetMinLength(1)
	hostField.SetMaxLength(255)
	if domain.ID() != "" {
		hostField.SetValue(domain.Host())
	}
	if err := validator.AddField(hostField); err != nil {
		return nil, errors.Trace(err)
	}

	languageField := validator.NewChoiceField("language")
	languageField.SetOptional()
	languageField.SetSearchable()
	languageField.SetCheckValuesInChoices()
	languageField.SetChoices(f.validatorsFactory.GetLanguageOptionsForLanguage(language))
	if domain.ID() != "" && domain.Language() != nil {
		languageField.SetValue(strconv.FormatUint(uint64(*domain.Language()), 10))
	}
	if err := validator.AddField(languageField); err != nil {
		return nil, errors.Trace(err)
	}

	currenciesField := validator.NewChoiceField("currencies")
	currenciesField.SetOptional()
	currenciesField.SetSearchable()
	currenciesField.SetCheckValuesInChoices()
	currenciesField.SetMultiple()
	currenciesField.SetPlaceholder(f.translationsRepository.Plural(language.ID(), "currency") + " (" +
		f.translationsRepository.Singular(language.ID(), "optional") + ")")
	currenciesField.SetChoices(f.validatorsFactory.GetCurrencyOptionsForLanguage(language))
	if domain.ID() != "" {
		if err := currenciesField.SetValues(strings.Split(domain.Currencies(), ",")); err != nil {
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

	validator.SetPreViewCallback(func(v validators.Validator) {
		siteID := siteField.Value()
		if siteID == "" {
			siteID = domain.SiteID()
		}
		if siteID != "" {
			site, ok, err := f.siteStore.GetOne(siteID)
			if err != nil {
				panic(errors.ErrorStack(err))
			}
			if !ok {
				panic(errors.Errorf("no site found for ID `%s` when one should exist", siteID))
			}
			siteField.AddChoice(f.validatorsFactory.NewChoiceOption(siteID, f.siteStore.GetTranslatedName(site, language.ID())))
			siteField.SetValue(siteID)
		}
	})

	return &CreateUpdate{
		validator: validator,
		view:      formview.New(validator),
	}, nil
}

// New returns a new instance of LoginForm.
func New(validatorsFactory validators.Factory, translationsRepository translations.Repository,
	configService config.Config, siteStore site.Store) *Forms {
	return &Forms{
		validatorsFactory:      validatorsFactory,
		translationsRepository: translationsRepository,
		configService:          configService,
		siteStore:              siteStore,
	}
}
