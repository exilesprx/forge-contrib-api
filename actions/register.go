package actions

import (
	"errors"
	"net/http"

	"github.com/exilesprx/forge-contrib-api/database"
	"github.com/exilesprx/forge-contrib-api/models"
	"github.com/exilesprx/forge-contrib-api/values"
	"github.com/labstack/echo"
)

// Register action to register a UnregisteredUser
type Register struct {
	Repository database.UserRepository
}

// RegisterUser registers a new user
func (register *Register) RegisterUser(context echo.Context) error {

	unregisteredUser := models.UnregisteredUser{}

	context.Bind(&unregisteredUser)

	err := register.registerUser(unregisteredUser)

	if err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	return context.JSON(http.StatusOK, "User registered")
}

func (register *Register) registerUser(user models.UnregisteredUser) error {

	if register.userExists(user) {
		return errors.New("User exists")
	}

	return register.Repository.CreateUser(user)
}

func (register *Register) userExists(user models.UnregisteredUser) bool {

	email := values.Email{
		Address: user.Email,
	}

	_, err := register.Repository.FindUserByEmail(email)

	if err != nil {
		return false
	}

	return true
}
