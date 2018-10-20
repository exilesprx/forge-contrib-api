package main

import (
	"os"

	"github.com/exilesprx/forge-contrib-api/actions"
	"github.com/exilesprx/forge-contrib-api/database"

	"github.com/labstack/echo"
)

func main() {
	app := echo.New()

	connection := database.CreateConnection()

	registerRoutes(app, connection)

	startApplication(app)
}

func registerRoutes(app *echo.Echo, connection database.Connection) {

	register := actions.Register{
		Repository: database.UserRepository{
			Connection: connection,
		},
	}

	app.POST("/register", register.RegisterUser)

	app.POST("/login", actions.Login)
}

func startApplication(app *echo.Echo) {

	port := os.Getenv("ECHO_PORT")

	err := app.Start(":" + port)

	app.Logger.Fatal(err)
}
