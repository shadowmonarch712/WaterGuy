package main

import (
    "context"
    "log"
    "gofr.dev/pkg/gofr"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
)

type UserEntry struct {
    UserID string `json:"userID"`
    Value  int    `json:"value"`
}

func main() {
    // initialise gofr object
    app := gofr.New()

    // MongoDB connection string
    mongoURI := "mongodb+srv://shadowmonarch712:miniproj@cluster0.x6gqsuf.mongodb.net/"


    // Connect to MongoDB
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
    if err != nil {
        log.Fatal(err)
    }

    // Check the connection
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    } else {
        log.Println("Connected to MongoDB!")
    }

    // collection configuration
    collection := client.Database("your_database").Collection("your_collection")

    // Create user 
    app.POST("/store", func(ctx *gofr.Context) (interface{}, error) {
        var userEntry UserEntry

        // Decode the request body into userEntry
        if err := ctx.Bind(&userEntry); err != nil {
            return nil, err
        }

        // Create a BSON document from the userEntry
        doc := bson.D{{"userID", userEntry.UserID}, {"value", userEntry.Value}}

        // document insertion
        insertResult, err := collection.InsertOne(context.TODO(), doc)
        if err != nil {
            return nil, err
        }

        return insertResult.InsertedID, nil
    })
    app.GET("/fetch", func(ctx *gofr.Context) (interface{}, error) {

        // Get the userID from the query parameters
        params := ctx.Params() 
        userID := params["UserID"] 
    
        // filter to search for the document
        filter := bson.D{{"userID", userID}}
    
        // Find the document in the collection
        var result UserEntry
        err := collection.FindOne(context.TODO(), filter).Decode(&result)

        if err != nil {
            if err == mongo.ErrNoDocuments {
                return "No document found with the given userID", nil
            }
            return nil, err
        }
    
        return result, nil
    })
    
    app.PUT("/update", func(ctx *gofr.Context) (interface{}, error) {
        // Get the userID and new value from the request body
        var userEntry UserEntry
        if err := ctx.Bind(&userEntry); err != nil {
            return nil, err
        }
    
        // filter to search for the document
        filter := bson.D{{"userID", userEntry.UserID}}
    
        // Create an update to change the value
        update := bson.D{{"$set", bson.D{{"value", userEntry.Value}}}}
    
        // Update the document in the collection
        updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
        if err != nil {
            return nil, err
        }
    
        // Check if the document was found and updated
        if updateResult.MatchedCount == 0 {
            return "No document found with the given userID to update", nil
        }
    
        return updateResult.ModifiedCount, nil
    })
    
    app.DELETE("/delete", func(ctx *gofr.Context) (interface{}, error) {
        // Get the userID from the query parameters
        userID := ctx.Params()["userID"]
    
        // filter to search for the document
        filter := bson.D{{"userID", userID}}
    
        // Delete the document from the collection
        deleteResult, err := collection.DeleteOne(context.TODO(), filter)
        if err != nil {
            return nil, err
        }
    
        // Check if the document was found and deleted
        if deleteResult.DeletedCount == 0 {
            return "No document found with the given userID to delete", nil
        }
    
        return deleteResult.DeletedCount, nil
    })
    // app will start on port 8000
    app.Start()

    // connection close
    defer func() {
        if err = client.Disconnect(context.TODO()); err != nil {
            log.Fatal(err)
        }
    }()
}
 