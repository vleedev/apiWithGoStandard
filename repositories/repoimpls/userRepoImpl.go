package repoimpls

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"time"
	"vlee/databases"
	"vlee/models"
	"vlee/repositories"
)

type UserRepoImpl struct {
	MongoCollection	*mongo.Collection
}
func NewUserRepo() repositories.UserRepo {
	return &UserRepoImpl{MongoCollection:databases.DBSessions.MongoInstance.Collection("Users")}
}
func (u UserRepoImpl) AllUsers() ([]*models.User, error) {
	users := make([]*models.User, 0)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := u.MongoCollection.Find(ctx, bson.D{})
	if err != nil {
		return users, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var user *models.User
		err := cur.Decode(&user)
		if err != nil {
			return users, err
		}
		// do something with result....
		users = append(users, user)
	}
	if err := cur.Err(); err != nil {
		return users, err
	}
	return users, nil
}
func (u UserRepoImpl) InsertOneUser(user *models.User) error {
	user.CreatedAt = time.Now().UnixNano()
	user.UpdatedAt = time.Now().UnixNano()
	_, err := u.MongoCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}
func (u UserRepoImpl) FindUserByEmail(email *string) (*models.User, error) {
	var user *models.User
	err := u.MongoCollection.FindOne(context.TODO(), bson.D{{"email", email}}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}
func (u UserRepoImpl) CheckSignInInfo(email,password *string) (*models.User, error) {
	var user *models.User
	err := u.MongoCollection.FindOne(context.TODO(), bson.D{{"email", email}}).Decode(&user)
	if err != nil {
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(*password))
	if err != nil {
		return user, err
	}
	return user, err
}
func (u UserRepoImpl) GetUserByObjectID(id *primitive.ObjectID) (*models.User, error) {
	var user *models.User
	err := u.MongoCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}
func (u UserRepoImpl) CheckUserExistence(email *string) (bool, error) {
	var user *models.User
	err := u.MongoCollection.FindOne(context.TODO(), bson.D{{"email", email}}).Decode(&user)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return false, nil
		}
		return true, err
	}
	return true, nil
}