package main

import (
 "github.com/labstack/echo/v4"
 "github.com/labstack/echo/v4/middleware"
 "net/http"
 "os"
)

func getGo(c echo.Context) error {
	hostname, err := os.Hostname()

	if err != nil {
		return c.String(http.StatusOK, "Hello, World!\nError:"+err.Error())
	}

	return c.String(http.StatusOK, "Hello, World!\nHostname:"+hostname)
}

func getSay(c echo.Context) error {
	hostname, err := os.Hostname()

	if err != nil {
		return c.String(http.StatusOK, "Привет, мир!\nОшибка:"+err.Error())
	}
	
	return c.String(http.StatusOK, "Привет, мир!\nНаименование хоста:"+hostname)
}

func main() {
 e := echo.New()
 
 // Middleware
 e.Use(middleware.Logger())
 e.Use(middleware.Recover())
 
 // Routes
 e.GET("/go", getGo)
 e.GET("/say", getSay)
 
 // Start server
 e.Logger.Fatal(e.Start(":8079"))
}
