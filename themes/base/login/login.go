package login

import (
	"io"

	"github.com/espal-digital-development/espal-core/repositories/themes"
)

// New returns a new instance of a Login View.
func New() themes.Viewable {
	view := themes.NewView("login")
	view.SetCallback(func(w io.Writer, data themes.DataStore) error {
		w.Write([]byte("Stub test"))
		v, ok := data.Get("test")
		if ok {
			w.Write([]byte(v.(string)))
		}
		return nil
	})
	return view
}
