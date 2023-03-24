package user

import (
	"context"
	"net/http"

	"github.com/eriickz/go-user-api/config"
	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	var req UserRequest

	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "You need to provide a valid user data.")
	}

	user := User{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Avatar:    req.Avatar,
	}

	if _, err := config.DB.NewInsert().Model(&user).Exec(context.TODO()); err != nil {
		return UserErrorHandler(err, c)
	}

	return c.JSON(http.StatusOK, user)
}

func GetUserById(c echo.Context) error {
	var req struct {
		Id int64 `json:"id" query:"id"`
	}

	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "You need to provide a valid user data.")
	}

	var user User

	if err := config.DB.NewSelect().Model(&user).Where("id = ?", req.Id).Scan(context.TODO()); err != nil {
		return UserErrorHandler(err, c)
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateUserById(c echo.Context) error {
	var req UserRequest

	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "You need to provide a valid user data.")
	}

	return c.String(http.StatusOK, "Testing!...")
}
