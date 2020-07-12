package createupdate

import (
	"strconv"

	"github.com/espal-digital-development/espal-core/config"
	usercontact "github.com/espal-digital-development/espal-core/stores/user/contact"
	"github.com/espal-digital-development/espal-core/validators"
	"github.com/espal-digital-development/espal-core/validators/formview"
	"github.com/juju/errors"
)

var _ Factory = &Forms{}

type language interface {
	ID() uint16
	Code() string
}

type contact interface {
	ID() string
	Comments() *string
	Sorting() uint
}

// Factory represents an object that serves new forms.
type Factory interface {
	New(contact contact, language language) (Form, error)
}

// Forms holds all logic to spawn forms for this type.
type Forms struct {
	validatorsFactory validators.Factory
	configService     config.Config
	userContactStore  usercontact.Store
}

// New creates a new Form instance with the required logic.
func (f *Forms) New(contact contact, language language) (Form, error) {
	validator, err := f.validatorsFactory.NewForm(language)
	if err != nil {
		return nil, errors.Trace(err)
	}

	commentsField := validator.NewTextAreaField("comments")
	commentsField.SetOptional()
	commentsField.SetTranslatePlural()
	if contact.ID() != "" && contact.Comments() != nil {
		commentsField.SetValue(*contact.Comments())
	}
	if err := validator.AddField(commentsField); err != nil {
		return nil, errors.Trace(err)
	}

	sortingField := validator.NewNumberField("sorting")
	sortingField.SetOptional()
	if contact.ID() != "" && contact.Sorting() > 0 {
		sortingField.SetValue(strconv.FormatUint(uint64(contact.Sorting()), 10))
	}
	if err := validator.AddField(sortingField); err != nil {
		return nil, errors.Trace(err)
	}

	contactField := validator.NewChoiceField("contact")
	contactField.SetSearchable()
	contactField.SetSearchableDataPath(f.configService.AdminURL() + "/User/Search")
	if err := validator.AddField(contactField); err != nil {
		return nil, errors.Trace(err)
	}

	validator.SetPreViewCallback(func(v validators.Validator) {
		contactID := contactField.Value()
		if contactID == "" {
			contactID = contact.ID()
		}
		if contactID != "" {
			contactUser, ok, err := f.userContactStore.GetOneByIDWithCreator(contactID)
			if err != nil {
				panic(errors.ErrorStack(err))
			}
			if !ok {
				contactField.AddChoice(f.validatorsFactory.NewChoiceOption(contactID,
					f.userContactStore.Name(contactUser, language.ID())))
				contactField.SetValue(contactID)
			}
		}
	})

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
func New(validatorsFactory validators.Factory, configService config.Config, userContactStore usercontact.Store) *Forms {
	return &Forms{
		validatorsFactory: validatorsFactory,
		configService:     configService,
		userContactStore:  userContactStore,
	}
}
