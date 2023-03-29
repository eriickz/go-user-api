package usertests

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"testing"

	"github.com/eriickz/go-user-api/config"
	"github.com/eriickz/go-user-api/modules/user"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var mockUser = user.User{
	Firstname: "Test",
	Lastname:  "Just a Test",
	Email:     "testing@example.com",
	Avatar:    "avatar.jpg",
}

func TestUsersEndpoints(t *testing.T) {
	config.ConnectAndLoadDB()
	t.Run("Create test user", createMockUser)
	t.Run("Get all users", getAllUsers)
	t.Run("Get test user", getMockUser)
	t.Run("Update test user", updateMockUser)
	t.Run("Delete test user", deleteMockUser)
}

func createMockUser(t *testing.T) {
	e := echo.New()
	context, recorder := CreateTestContextAndRecorder(e, t, mockUser, http.MethodPost, "/api/user/create")

	if assert.NoError(t, user.CreateUser(context)) {
		var testUser user.User

		if assert.Equal(t, http.StatusCreated, recorder.Code) {
			assert.NoError(t, json.Unmarshal([]byte(recorder.Body.String()), &testUser))
			assert.Equal(t, mockUser.Firstname, testUser.Firstname)
			assert.Equal(t, mockUser.Lastname, testUser.Lastname)
			assert.Equal(t, mockUser.Email, testUser.Email)
			assert.Equal(t, mockUser.Avatar, testUser.Avatar)
			mockUser = testUser
		}
	}
}

func getAllUsers(t *testing.T) {
	e := echo.New()
	context, recorder := CreateTestContextAndRecorder(e, t, nil, http.MethodGet, "/api/user/getUsers")

	if assert.NoError(t, user.GetUsers(context)) {
		var users []user.User

		if assert.Equal(t, http.StatusOK, recorder.Code) {
			assert.NoError(t, json.Unmarshal([]byte(recorder.Body.String()), &users))
			assert.True(t, SearchForMockUser(users, mockUser))
		}
	}
}

func getMockUser(t *testing.T) {
	e := echo.New()
	query := make(url.Values)
	query.Set("id", strconv.FormatInt(mockUser.ID, 10))

	context, recorder := CreateTestContextAndRecorder(e, t, nil, http.MethodGet, "/api/user/getUserById?"+query.Encode())

	if assert.NoError(t, user.GetUserById(context)) {
		var testUser user.User

		if assert.Equal(t, http.StatusOK, recorder.Code) {
			assert.NoError(t, json.Unmarshal([]byte(recorder.Body.String()), &testUser))
			assert.Equal(t, mockUser, testUser)
		}
	}
}

func updateMockUser(t *testing.T) {
	e := echo.New()
	mockUser.Lastname = "Testing"

	context, recorder := CreateTestContextAndRecorder(e, t, mockUser, http.MethodPut, "/api/user/updateUser")

	if assert.NoError(t, user.UpdateUser(context)) {
		var testUser user.User

		if assert.Equal(t, http.StatusOK, recorder.Code) {
			assert.NoError(t, json.Unmarshal([]byte(recorder.Body.String()), &testUser))
			assert.Equal(t, mockUser.Lastname, testUser.Lastname)
			mockUser = testUser
		}
	}
}

func deleteMockUser(t *testing.T) {
	e := echo.New()
	query := make(url.Values)
	query.Set("id", strconv.FormatInt(mockUser.ID, 10))

	context, recorder := CreateTestContextAndRecorder(e, t, mockUser, http.MethodDelete, "/api/user/deleteUserById?"+query.Encode())

	if assert.NoError(t, user.DeleteUserById(context)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
	}
}
