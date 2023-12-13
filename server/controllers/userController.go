package controllers

import (
    "context"
    "waterguy/database"
    "waterguy/models"
    "gofr.dev/pkg/gofr"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
)



func CreateUserEntry(ctx *gofr.Context) (interface{}, error) {
    var userEntry models.UserEntry
	collection := database.GetCollection("user_goal")
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
}

func FetchUserEntry(ctx *gofr.Context) (interface{}, error) {
	collection := database.GetCollection("user_goal")
	params := ctx.Params() 
	userID := params["UserID"] 

	// filter to search for the document
	filter := bson.D{{"userID", userID}}

	// Find the document in the collection
	var result models.UserEntry
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "No document found with the given userID", nil
		}
		return nil, err
	}

	return result, nil
}

func UpdateUserEntry(ctx *gofr.Context) (interface{}, error) {
	collection := database.GetCollection("user_goal")
	var userEntry models.UserEntry
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
}

func DeleteUserEntry(ctx *gofr.Context) (interface{}, error) {
	collection := database.GetCollection("user_goal")
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
}
