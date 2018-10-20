package models

type UnregisteredUser struct {
	Name     string `json:"name" query:"name" form:"name"`
	Email    string `json:"email" query:"email" form:"email"`
	Password string `json:"password" query:"password" form:"password"`
}
