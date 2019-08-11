package databases

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)
/*
*	The connect driver for mongodb
*	Author: Lee Tuan
 */
func ConnectMongoDB() {
	a	:= os.Getenv("DB_ADDRESS")
	p 	:= os.Getenv("DB_PORT")
	d	:= os.Getenv("DB_DATABASE")
	u	:= os.Getenv("DB_USERNAME")
	pass	:= os.Getenv("DB_PASSWORD")
	// Set client options
	connectStr := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",u,pass,a,p,d)
	clientOptions := options.Client().ApplyURI(connectStr)
	// Connect to MongoDB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connect to mongodb on %s successfully", a)
	DBSessions.MongoInstance = client.Database(d)
}