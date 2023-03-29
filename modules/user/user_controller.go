package user

import (
	"context"
	"net/http"

	"github.com/eriickz/go-user-api/config"
	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	var req UserRequest

	BindUserRequest(&req, c)

	user := User{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Avatar:    req.Avatar,
	}

	if _, err := config.DB.NewInsert().Model(&user).Exec(context.TODO()); err != nil {
		return UserDBErrorHandler(err, c)
	}

	return c.JSON(http.StatusCreated, user)
}

func GetUsers(c echo.Context) error {
	var users []User

	if err := config.DB.NewSelect().Model(&users).OrderExpr("id DESC").Scan(context.TODO()); err != nil {
		return UserDBErrorHandler(err, c)
	}

	return c.JSON(http.StatusOK, users)
}

func GetUserById(c echo.Context) error {
	var req struct {
		Id int64 `query:"id"`
	}

	BindUserRequest(&req, c)

	var user User

	if err := config.DB.NewSelect().Model(&user).Where("id = ?", req.Id).Scan(context.TODO()); err != nil {
		return UserDBErrorHandler(err, c)
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	var req UserRequest

	BindUserRequest(&req, c)

	if req.Id == 0 {
		return c.String(http.StatusBadRequest, "You need to provide a valid user id.")
	}

	user := User{
		ID:        req.Id,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Avatar:    req.Avatar,
	}

	res, err := config.DB.NewUpdate().Model(&user).OmitZero().WherePK().Returning("*").Exec(context.TODO())

	if err != nil {
		return UserDBErrorHandler(err, c)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return UserDBErrorHandler(err, c)
	}

	if rowsAffected == 0 {
		return c.String(http.StatusBadRequest, "User not found.")
	}

	return c.JSON(http.StatusOK, user)
}

func DeleteUserById(c echo.Context) error {
	var req struct {
		Id int64 `query:"id"`
	}

	BindUserRequest(&req, c)

	if _, err := config.DB.NewDelete().Model(&User{ID: req.Id}).WherePK().Exec(context.TODO()); err != nil {
		return UserDBErrorHandler(err, c)
	}

	return c.String(http.StatusOK, "The user has been deleted")
}
