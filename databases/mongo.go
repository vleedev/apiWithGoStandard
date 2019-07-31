package databases

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func MongoDB() *mongo.Database {
	a	:= os.Getenv("DB_ADDRESS")
	p 	:= os.Getenv("DB_PORT")
	d	:= os.Getenv("DB_DATABASE")
	u	:= os.Getenv("DB_USERNAME")
	pass	:= os.Getenv("DB_PASSWORD")
	// Set client options
	uri := "mongodb://" + u + ":" + pass + "@" + a + ":" + p + "/" + d
	clientOptions := options.Client().ApplyURI(uri)
	// Connect to MongoDB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println(err)
	}
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Connect to mongodb successfully")
	return client.Database(d)
}