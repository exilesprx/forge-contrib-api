package main

import (
	"github.com/exilesprx/forge-contrib-api/routes"

	"github.com/labstack/echo"
)

func main() {
	var app *echo.Echo = echo.New()

	app.POST("/login", routes.Login)

	var startResponse error = app.Start(":1323")

	app.Logger.Fatal(startResponse)
}
