package databases

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func MongoDB() (*mongo.Database, error) {
	a	:= os.Getenv("DB_ADDRESS")
	p 	:= os.Getenv("DB_PORT")
	d	:= os.Getenv("DB_DATABASE")
	u	:= os.Getenv("DB_USERNAME")
	pass	:= os.Getenv("DB_PASSWORD")
	// Set client options
	uri := "mongodb://" + u + ":" + pass + "@" + a + ":" + p + "/" + d
	clientOptions := options.Client().ApplyURI(uri)
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	return client.Database(d), nil
}