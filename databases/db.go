package databases

import "go.mongodb.org/mongo-driver/mongo"

type DatabaseSessions struct {
	MongoInstance *mongo.Database
}
var DBSessions DatabaseSessions