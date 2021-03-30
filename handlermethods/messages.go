package handlermethods

import (
	"context"
	"errors"
	"fmt"
	"time"

	"../models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MessageCollection = models.MessageCollection

func CreateMessage(message, email string) (string, error) {
	time := time.Now()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:9001"))
	if err != nil {
		fmt.Println("Error Connecting to DB", err)
		return "", errors.New("internal server error")
	}
	collection := client.Database("gotest").Collection("message")
	_, err = collection.InsertOne(context.TODO(), bson.M{"message": message, "createdBy": email, "createdDate": time})
	if err != nil {
		fmt.Println("Error Inserting document", err)
		return "", errors.New("internal server error")
	}
	return "Successfully Created Message", nil
}

func GetMessage(email string) ([]MessageCollection, error) {
	var messages []MessageCollection
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:9001"))
	if err != nil {
		fmt.Println("Error Connecting to DB", err)
		return messages, errors.New("internal server error")
	}
	collection := client.Database("gotest").Collection("message")
	cursor, err := collection.Find(ctx, bson.M{"createdBy": email})
	if err != nil {
		fmt.Println("Error Getting data", err)
		return messages, errors.New("internal server error")
	}
	for cursor.Next(ctx) {
		var message MessageCollection
		if err = cursor.Decode(&message); err != nil {
			fmt.Println("Error Converting data", err)
			return messages, errors.New("internal server error")
		}
		messages = append(messages, message)
	}
	return messages, nil
}

func DeleteMessage(messageId string) (string, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:9001"))
	if err != nil {
		fmt.Println("Error Connecting to DB", err)
		return "", errors.New("internal server error")
	}
	_id, err := primitive.ObjectIDFromHex(messageId)
	if err != nil {
		fmt.Println("Error Converting Id", err)
		return "", errors.New("internal server error")
	}
	collection := client.Database("gotest").Collection("message")
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": _id})
	if err != nil {
		fmt.Println("Error Deleting document", err)
		return "", errors.New("internal server error")
	}
	return "Successfully Deleted Message", nil
}
