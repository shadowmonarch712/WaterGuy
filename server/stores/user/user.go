package User

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gofr.dev/pkg/gofr"
	"waterguy/database"
	"waterguy/models"
)

type store struct{}

func New() store {
	return store{}
}
func (s store) CreateUserEntry(ctx *gofr.Context, c models.UserEntry) error {
	// fetch the Mongo collection
	collection := database.GetCollection("user_goal")

	_, err := collection.InsertOne(ctx, c)

	return err
}
func (s store) FetchUserEntry(ctx *gofr.Context, userID string) (interface{}, error) {
	// fetch the Mongo collection
	collection := database.GetCollection("user_goal")

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

func (s store) UpdateUserEntry(ctx *gofr.Context) (interface{}, error) {
	// fetch the Mongo collection
	collection := database.GetCollection("user_goal")
	var userEntry models.UserEntry
	if err := ctx.Bind(&userEntry); err != nil {
		return nil, err
	}

	// Filter to search for the document
	filter := bson.D{{"userID", userEntry.UserID}}

	// Create an update to change the value
	update := bson.D{{"$set", bson.D{{"value", userEntry.Value}}}}

	// Options to return the updated document
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	// Find one document and update it in the collection
	var updatedDocument models.UserEntry
	err := collection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedDocument)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "No document found with the given userID to update", nil
		}
		return nil, err
	}

	// Return the updated value field of the document
	return updatedDocument.Value, nil
}

func (s store) DeleteUserEntry(ctx *gofr.Context, userID string) (interface{}, error) {
	// fetch the Mongo collection
	collection := database.GetCollection("user_goal")
	// Get the userID from the query parameters
	//userID := ctx.Params()["userID"]

	// filter to search for the document
	filter := bson.D{{"userID", userID}}

	// Delete the document from the collection
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	// Check if the document was found and deleted
	if deleteResult.DeletedCount == 0 {
		return 0, nil
	}

	return int(deleteResult.DeletedCount), nil
}
