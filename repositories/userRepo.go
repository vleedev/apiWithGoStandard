package repositories

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"vlee/models"
)
type UserRepo interface {
	AllUsers() ([]*models.User, error)
	FindUserByEmail(email *string) (*models.User, error)
	InsertOneUser(u *models.User) error
	CheckSignInInfo(email,password *string) (*models.User, error)
	GetUserByObjectID(id *primitive.ObjectID) (*models.User, error)
	CheckUserExistence(email *string) (bool, error)
}
