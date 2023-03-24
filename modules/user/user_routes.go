package user

import (
	"context"

	"github.com/eriickz/go-user-api/config"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(apiGroup *echo.Group) {
	config.DB.NewCreateTable().Model(&User{}).Exec(context.TODO())
	userGroup := apiGroup.Group("/user")

	userGroup.GET("/getUserById", GetUserById)
	userGroup.POST("/create", CreateUser)
	userGroup.PUT("/updateUser", UpdateUser)
	userGroup.DELETE("/deleteUserById", DeleteUserById)
}
