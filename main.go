package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func getRoot(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getSay(c echo.Context) error {
	return c.String(http.StatusOK, "Привет, Мир!")
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", getRoot)
	e.GET("/say", getSay)

	// Start server
	e.Logger.Fatal(e.Start(":8079"))
}
