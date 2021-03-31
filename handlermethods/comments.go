package handlermethods

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ritwik1233/go-rest-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CommentsCollection = models.CommentsCollection

func CreateComment(comment, messageId, createdBy string) (string, error) {
	time := time.Now()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGOURI")))
	if err != nil {
		fmt.Println("Error Connecting to DB", err)
		return "", errors.New("internal server error")
	}
	_id, err := primitive.ObjectIDFromHex(messageId)
	if err != nil {
		fmt.Println("Error Converting Id", err)
		return "", errors.New("internal server error")
	}
	collection := client.Database("gotest").Collection("comments")
	_, err = collection.InsertOne(context.TODO(), bson.M{"comment": comment, "messageId": _id, "createdBy": createdBy, "createdDate": time})
	if err != nil {
		fmt.Println("Error Inserting document", err)
		return "", errors.New("internal server error")
	}
	return "Successfully Created Comment", nil
}

func GetComments(messageId string) ([]CommentsCollection, error) {
	var comments []CommentsCollection
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOURI")))
	if err != nil {
		fmt.Println("Error Connecting to DB", err)
		return comments, errors.New("internal server error")
	}
	collection := client.Database("gotest").Collection("comments")
	_id, err := primitive.ObjectIDFromHex(messageId)
	if err != nil {
		fmt.Println("Error Converting Id", err)
		return comments, errors.New("internal server error")
	}
	cursor, err := collection.Find(ctx, bson.M{"messageId": _id})
	if err != nil {
		fmt.Println("Error Getting data", err)
		return comments, errors.New("internal server error")
	}
	for cursor.Next(ctx) {
		var comment CommentsCollection
		if err = cursor.Decode(&comment); err != nil {
			fmt.Println("Error Converting data", err)
			return comments, errors.New("internal server error")
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func DeleteComment(commentId string) (string, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOURI")))
	if err != nil {
		fmt.Println("Error Connecting to DB", err)
		return "", errors.New("internal server error")
	}
	_id, err := primitive.ObjectIDFromHex(commentId)
	if err != nil {
		fmt.Println("Error Converting Id", err)
		return "", errors.New("internal server error")
	}
	collection := client.Database("gotest").Collection("comments")
	_, err = collection.DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		fmt.Println("Error Deleting Comments", err)
		return "", errors.New("internal server error")
	}
	return "Successfully Deleted Comment", nil
}

func UpdateComment(commentId, updatedComment string) (string, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOURI")))
	if err != nil {
		fmt.Println("Error Connecting to DB", err)
		return "", errors.New("internal server error")
	}
	_id, err := primitive.ObjectIDFromHex(commentId)
	if err != nil {
		fmt.Println("Error Converting Id", err)
		return "", errors.New("internal server error")
	}
	collection := client.Database("gotest").Collection("comments")
	_, err = collection.UpdateOne(
		ctx,
		bson.M{"_id": _id},
		bson.M{
			"$set": bson.M{"comment": updatedComment},
		},
	)
	if err != nil {
		fmt.Println("Error Deleting document", err)
		return "", errors.New("internal server error")
	}
	return "Successfully Updated Comment", nil
}
