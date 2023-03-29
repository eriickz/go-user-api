package usertests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/eriickz/go-user-api/modules/user"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

func CreateTestContextAndRecorder(e *echo.Echo, t *testing.T, mockData interface{}, method string, url string) (echo.Context, *httptest.ResponseRecorder) {
	body, err := json.Marshal(mockData)

	if assert.NoError(t, err) {
		var req *http.Request

		req = httptest.NewRequest(method, url, nil)

		if mockData != nil {
			req = httptest.NewRequest(method, url, strings.NewReader(string(body)))
		}

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		return e.NewContext(req, rec), rec
	}

	if err != nil {
		t.Fatal("Cannot create test context.")
	}

	return nil, nil
}

func SearchForMockUser(users []user.User, mockUser user.User) bool {
	foundIndex := slices.IndexFunc(users, func(user user.User) bool {
		return user.Firstname == mockUser.Firstname
	})

	return foundIndex > -1
}
