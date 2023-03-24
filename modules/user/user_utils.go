package user

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func UserDBErrorHandler(err error, c echo.Context) error {
	errorString := err.Error()

	switch true {
	case strings.Contains(errorString, "no rows in result set"):
		return c.String(http.StatusNotFound, "No user(s) found.")

	case strings.Contains(errorString, "duplicate") && strings.Contains(errorString, "email"):
		return c.String(http.StatusBadRequest, "The email are in use.")

	default:
		return c.String(http.StatusInternalServerError, "An error occurred.")
	}
}

func BindUserRequest(req interface{}, c echo.Context) interface{} {
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "Invalid user data supplied.")
	}

	return req
}
