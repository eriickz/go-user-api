package main

import (
	"github.com/eriickz/go-user-api/config"
	"github.com/eriickz/go-user-api/modules/user"
	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectAndLoadDB()
	e := echo.New()

	apiGroup := e.Group("/api")
	user.RegisterRoutes(apiGroup)

	e.Logger.Fatal(e.Start(":5000"))
}
