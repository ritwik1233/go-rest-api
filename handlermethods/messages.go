package handlermethods

import (
	"context"
	"fmt"
	"time"

	"../models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MessageCollection = models.MessageCollection

// create message
func CreateMessage(message, email string) (string, error) {
	time := time.Now()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:9001"))
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
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:9001"))
	if err != nil {
		fmt.Println("Error Connecting to DB", err)
		return messages, nil
	}
	collection := client.Database("gotest").Collection("message")
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
func DeleteMessage(messageId string) (string, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:9001"))
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
