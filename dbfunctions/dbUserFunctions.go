package dbfunctions

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/******User Functions Functions START****/
type UserCollection struct {
	ID       string `bson:"_id"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Username string `bson:"username"`
}

func checkIfUserExists(email string) (UserCollection, error) {
	var result UserCollection
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:9001"))
	if err != nil {
		return result, err
	}
	collection := client.Database("gotest").Collection("users")
	filter := bson.M{"email": email}
	ctx, cancel = context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		noDocumentMessage := "mongo: no documents in result"
		if err.Error() != noDocumentMessage {
			fmt.Println("Error Checking document", err)
			return result, err
		}
	}
	return result, nil
}

func checkLoginCredentials(email, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:9001"))
	if err != nil {
		return "", err
	}
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	collection := client.Database("gotest").Collection("users")
	var result UserCollection
	filter := bson.M{"email": email, "password": password}
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		fmt.Println("Error Decoding Document", err)
		return "", err
	}
	return "Login Successfull", nil
}

func registerUser(email, username, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:9001"))
	if err != nil {
		return "", err
	}
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	collection := client.Database("gotest").Collection("users")
	userDetails, err := checkIfUserExists(email)
	if err != nil {
		fmt.Println("Error Checking document", err)
		return "", err
	}
	if userDetails.Email == email {
		fmt.Println("User Already Exists")
		return "", nil
	}
	res, err := collection.InsertOne(ctx, bson.M{"email": email, "username": username, "password": password})
	if err != nil {
		fmt.Println("Error Inserting document", err)
		return "", err
	}
	fmt.Println("Successful Registered user", res)
	return "Successfully Registered user: " + userDetails.Email, nil
}
