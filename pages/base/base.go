package base

import (
	"io"

	"github.com/espal-digital-development/espal-core/sessions"
)

// Context core context.
type Context interface {
	io.Writer
	IsDevelopment() bool
	IsLoggedIn() bool
	HasAdminAccess() bool
	HasPprofEnabled() bool
	AdminURL() string
	PprofURL() string
	Translate(string) string
	TranslatePlural(string) string
	HasFlashMessage() bool
	GetFlashMessage() sessions.Message
}

// Form interactive handling.
type Form interface {
	Errors() string
	Open() string
	Field(string) string
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

// HasAdminAccess returns an indicator if the user has administrator access.
func (p *Page) HasAdminAccess() bool {
	return p.coreContext.HasAdminAccess()
}

// HasPprofEnabled returns an indicator if the user has pprof access.
func (p *Page) HasPprofEnabled() bool {
	return p.coreContext.HasPprofEnabled()
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

// PprofURL returns the url prefix for visiting pprof area paths.
func (p *Page) PprofURL() string {
	return p.coreContext.PprofURL()
}

// HasFlashMessage returns an indicator if a flash message was set.
func (p *Page) HasFlashMessage() bool {
	return p.coreContext.HasFlashMessage()
}

// GetFlashMessage returns the set flash message.
func (p *Page) GetFlashMessage() sessions.Message {
	return p.coreContext.GetFlashMessage()
}
