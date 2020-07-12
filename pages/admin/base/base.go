package base

import (
	"io"

	"github.com/espal-digital-development/espal-core/adminmenu"
	"github.com/espal-digital-development/espal-core/sessions"
)

// Context core context.
type Context interface {
	io.Writer
	IsDevelopment() bool
	IsLoggedIn() bool
	AdminURL() string
	Translate(string) string
	TranslatePlural(string) string
	HasFlashMessage() bool
	GetFlashMessage() sessions.Message
	HasUserRight(string) bool
	AdminMainMenu() []*adminmenu.Block
}

// Form interactive handling.
type Form interface {
	Errors() string
	Open() string
	Field(string) string
	CreateUpdateActions(string, string) string
	ContainsSelectSearch() bool
}

// Page template object.
type Page struct {
	coreContext Context
}

// SetCoreContext sets the basic context requirements of the p.
func (p *Page) SetCoreContext(context Context) {
	p.coreContext = context
}

// GetCoreContext returns the internal core context.
func (p *Page) GetCoreContext() Context {
	return p.coreContext
}

// IsDevelopment returns an indicator if the project is in development mode.
func (p *Page) IsDevelopment() bool {
	return p.coreContext.IsDevelopment()
}

// IsLoggedIn returns an indicator if the user is logged in or not.
func (p *Page) IsLoggedIn() bool {
	return p.coreContext.IsLoggedIn()
}

// Translate translates the given key based on the language
// active in the current context.
func (p *Page) Translate(key string) string {
	return p.coreContext.Translate(key)
}

// TranslatePlural translates the given key based on the language
// active in the current context in plural.
func (p *Page) TranslatePlural(key string) string {
	return p.coreContext.TranslatePlural(key)
}

// AdminURL returns the url prefix for visiting admin area paths.
func (p *Page) AdminURL() string {
	return p.coreContext.AdminURL()
}

// HasFlashMessage returns an indicator if a flash message was set.
func (p *Page) HasFlashMessage() bool {
	return p.coreContext.HasFlashMessage()
}

// GetFlashMessage returns the set flash message.
func (p *Page) GetFlashMessage() sessions.Message {
	return p.coreContext.GetFlashMessage()
}

// HasUserRight returns if the current user (if logged in) has the userright.
func (p *Page) HasUserRight(userRight string) bool {
	return p.coreContext.HasUserRight(userRight)
}

// AdminMainMenu returns the rendered admin menu for the current user.
func (p *Page) AdminMainMenu() []*adminmenu.Block {
	return p.coreContext.AdminMainMenu()
}
