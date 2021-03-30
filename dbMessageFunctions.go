package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/******Message Functions Functions START****/
type MessageCollection struct {
	ID          string    `bson:"_id"`
	Message     string    `bson:"message"`
	CreatedBy   string    `bson:"createdBy"`
	CreatedDate time.Time `bson:"createdDate"`
}

// create message
func CreateMessage(message, email string) (string, error) {
	time := time.Now()
	client, err := ConnectDB()
	if err != nil {
		fmt.Println("Error Connecting to DB", err)
		return "", nil
	}
	collection := client.Database("gotest").Collection("message")
	res, err := collection.InsertOne(context.TODO(), bson.M{"message": message, "createdBy": email, "createdDate": time})
	if err != nil {
		fmt.Println("Error Inserting document", err)
		return "", err
	}
	fmt.Println("Successful Created Message", res.InsertedID)
	return "Successfully Created Message", nil
}

// get message
func GetMessage(email string) ([]MessageCollection, error) {
	var messages []MessageCollection
	client, err := ConnectDB()
	if err != nil {
		fmt.Println("Error Connecting to DB", err)
		return messages, err
	}
	collection := client.Database("gotest").Collection("message")
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{"createdBy": email})
	if err != nil {
		fmt.Println("Error Getting data", err)
		return messages, err
	}
	for cursor.Next(ctx) {
		var message MessageCollection
		if err = cursor.Decode(&message); err != nil {
			fmt.Println("Error Converting data", err)
			return messages, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}

// delete message
func deleteMessage(messageId string) (string, error) {
	client, err := ConnectDB()
	if err != nil {
		fmt.Println("Error Connecting to DB", err)
		return "", nil
	}
	_id, err := primitive.ObjectIDFromHex(messageId)
	if err != nil {
		fmt.Println("Error Converting Id", err)
		return "", nil
	}
	collection := client.Database("gotest").Collection("message")
	res, err := collection.DeleteOne(context.TODO(), bson.M{"_id": _id})
	if err != nil {
		fmt.Println("Error Deleting document", err)
		return "", err
	}
	fmt.Println("Successfully Deleted Message", res)
	return "Successfully Deleted Message", nil
}
