package models

import (
	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

// RegisteredUser
type RegisteredUser struct {
	ID       objectid.ObjectID `bson:"_id" json:"id"`
	Name     string            `bson:"name" json:"name"`
	Email    string            `bson:"email" json:"email"`
	Password string            `bson:"password" json:"password"`
}
