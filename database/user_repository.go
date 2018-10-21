package database

import (
	"context"
	"errors"

	"github.com/mongodb/mongo-go-driver/mongo"

	"github.com/exilesprx/forge-contrib-api/models"
	"github.com/exilesprx/forge-contrib-api/values"
	"github.com/mongodb/mongo-go-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

const collectionName = "user"

// UserRepository user specific database operations
type UserRepository struct {
	Connection
}

// CreateUser Creates a new user in the database
func (repo *UserRepository) CreateUser(user models.UnregisteredUser) error {
	collection := getCollection(repo)

	encrytedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 1)

	record := bson.NewDocument(
		bson.EC.String("name", user.Name),
		bson.EC.String("email", user.Email),
		bson.EC.String("password", string(encrytedPassword)),
	)

	_, err := collection.InsertOne(context.Background(), record)

	if err != nil {
		return errors.New("Cannot create user")
	}

	return nil
}

// FindUserByEmail find a registered user by email address
func (repo *UserRepository) FindUserByEmail(email values.Email) (models.RegisteredUser, error) {
	collection := getCollection(repo)

	user := models.RegisteredUser{}

	filter := bson.NewDocument(bson.EC.String("email", email.Address))

	err := collection.FindOne(context.Background(), filter).Decode(&user)

	if err != nil {
		return user, errors.New("User not found")
	}

	return user, nil
}

func getCollection(repo *UserRepository) *mongo.Collection {
	return repo.Connection.Collection(collectionName)
}
