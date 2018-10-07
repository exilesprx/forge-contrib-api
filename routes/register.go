package routes

import (
	"net/http"

	"github.com/exilesprx/forge-contrib-api/database"
	"github.com/exilesprx/forge-contrib-api/models"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

// Register registers a new user
func Register(context echo.Context) error {

	database.Connect()

	var name string = context.FormValue("name")
	var password string = bcrypt.GenerateFromPassword(context.FormValue("password"))
	var email string context.FormValue("email")

	var user *models.User = new(models.User){name: name, email: email, password: password}

	database.CreateUser(user)

	return context.JSON(http.StatusOK, "Register")
}
