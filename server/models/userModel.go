package models

type UserEntry struct {
    UserID string `bson:"userID"`
    Value  int    `bson:"value"`
}
