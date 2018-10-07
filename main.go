package main

import (
	"github.com/exilesprx/forge-contrib-api/routes"

	"github.com/labstack/echo"
)

var port string = "1323"

func main() {
	var app *echo.Echo = echo.New()

	registerRoutes(app)

	startApplication(app)
}

func registerRoutes(app *echo.Echo) {
	app.POST("/register", routes.Register)

	app.POST("/login", routes.Login)
}

func startApplication(app *echo.Echo) {
	var err error = app.Start(":1323")

	app.Logger.Fatal(err)
}
