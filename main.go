package main

import (
	"net/http"
	"os"

	"github.com/eriickz/go-user-api/config"
	"github.com/eriickz/go-user-api/modules/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.ConnectAndLoadDB()
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodOptions, http.MethodDelete, http.MethodHead},
	}))

	apiGroup := e.Group("/api")
	user.RegisterRoutes(apiGroup)

	port := ":5000"

	if os.Getenv("ENVIRONMENT") != "testing" {
		port = ":8080"
	}

	e.Logger.Fatal(e.Start(port))
}
