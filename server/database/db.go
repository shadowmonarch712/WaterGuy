package database

import (
    "context"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func InitDB() {
    mongoURI := "mongodb://testAdmin:123@localhost:27017/?authMechanism=SCRAM-SHA-1"


    // Connect to MongoDB
	var err error
    Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
    if err != nil {
        log.Fatal(err)
    }

    // Check the connection
    err = Client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    } else {
        log.Println("Connected to MongoDB!")
    }

}
func GetCollection(collectionName string) *mongo.Collection {
    if Client == nil {
        log.Fatal("MongoDB client is not initialized")
    }
    return Client.Database("waterguy").Collection(collectionName)
}
