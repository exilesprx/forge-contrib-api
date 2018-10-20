package actions

import (
	"net/http"

	"github.com/labstack/echo"
)

// Login allow a user to login
func Login(context echo.Context) error {
	// TODO: Login a user in
	// TODO: Return a JWT token
	return context.JSON(http.StatusOK, "Login!")
}
