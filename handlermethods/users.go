package handlermethods

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/ritwik1233/go-rest-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserCollection = models.UserCollection

func GetUser(email string) (UserCollection, error) {
	var result UserCollection
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOURI")))
	if err != nil {
		return result, errors.New("internal server error")
	}
	collection := client.Database("gotest").Collection("users")
	filter := bson.M{"email": email}
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		noDocumentMessage := "mongo: no documents in result"
		if err.Error() != noDocumentMessage {
			fmt.Println("Error Checking document", err)
			return result, errors.New("internal server error")
		}
	}
	return result, nil
}

func CheckLoginCredentials(email, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOURI")))
	if err != nil {
		return "", errors.New("internal server error")
	}
	collection := client.Database("gotest").Collection("users")
	var result UserCollection
	filter := bson.M{"email": email, "password": password}
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		noDocumentMessage := "mongo: no documents in result"
		if err.Error() != noDocumentMessage {
			fmt.Println("Error Checking document", err)
			return "", errors.New("internal server error")
		} else {
			return "", errors.New("login failed: no user exists")
		}
	}
	return "Login Successfull", nil
}

func RegisterUser(email, username, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOURI")))
	if err != nil {
		return "", errors.New("internal server error")
	}
	collection := client.Database("gotest").Collection("users")
	userDetails, err := GetUser(email)
	if err != nil {
		fmt.Println("Error Checking document", err)
		return "", errors.New("internal server error")
	}
	if userDetails.Email == email {
		fmt.Println("User Already Exists")
		return "", errors.New("user already exists")
	}
	res, err := collection.InsertOne(ctx, bson.M{"email": email, "username": username, "password": password})
	if err != nil {
		fmt.Println("Error Inserting document", err)
		return "", errors.New("internal server error")
	}
	fmt.Println("Successful Registered user", res)
	return "Successfully Registered user: " + userDetails.Email, nil
}
