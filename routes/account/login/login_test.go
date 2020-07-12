package login_test

import (
	"testing"

	"github.com/espal-digital-development/espal-module-core/routes/account/login"
)

// TODO :: 777 Test fake a POST on the page with form filled validly (and a valid token)

// func getNewMockContext() contexts.Context {
// 	return &contextsmock.ContextMock{}
// }

// func getNewMockForm() base.Form {
// 	return &basemock.FormMock{}
// }

func TestNew(t *testing.T) {
	login := login.New(nil, nil)
	if login == nil {
		t.Fatal("result should not be nil")
	}
}

// TODO :: 777 Need to make this work more smooth. The page and form need to be mocked
// func TestHandle(t *testing.T) {
// 	validatorsFactory := nil
// 	userStore := nil
// 	page := page.New()
// 	form := form.New(validatorsFactory, userStore)

// 	login := New(page, nil)
// 	login.Handle(getNewMockContext())
// }
