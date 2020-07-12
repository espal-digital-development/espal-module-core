package login_test

import (
	"testing"

	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/routing/router/contexts/contextsmock"
	"github.com/espal-digital-development/espal-module-core/pages/account/login"
)

// TODO :: 777 The variation of Stream/Write/~normal~ causes misleading code coverage
// TODO :: 777 Need to finish the MockForm to make the test succeed again

func getNewMockContext() contexts.Context {
	return &contextsmock.ContextMock{}
}

// func getNewMockForm() base.Form {
// 	return &basemock.FormMock{}
// }

func newPage() login.Template {
	login := login.New()
	context := getNewMockContext()
	page := login.NewPage(context)
	return page
}

func TestNew(t *testing.T) {
	login := login.New()
	if login == nil {
		t.Fatal("result should not be nil")
	}
}

func TestNewPage(t *testing.T) {
	page := newPage()
	if page == nil {
		t.Fatal("result should not be nil")
	}
}

// func TestPageRender(t *testing.T) {
// 	page := newPage()
// 	p.Render()
// }
