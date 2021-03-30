package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
	client, err := ConnectDB()
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	collection := client.Database("gotest").Collection("users")
	filter := bson.M{"email": email}
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
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
	client, err := ConnectDB()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	collection := client.Database("gotest").Collection("users")
	var result UserCollection
	filter := bson.M{"email": email, "password": password}
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		fmt.Println("Error Decoding Document", err)
		return "", err
	}
	return "Login Successfull", nil
}

func registerUser(email, username, password string) string {
	client, err := ConnectDB()
	if err != nil {
		fmt.Println(err)
		return "Error Connecting to DB"
	}
	collection := client.Database("gotest").Collection("users")
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	userDetails, err := checkIfUserExists(email)
	if err != nil {
		fmt.Println("Error Checking document", err)
		return "Internal Server Error"
	}
	if userDetails.Email == email {
		fmt.Println("User Already Exists")
		return "User Already Exists"
	}
	res, err := collection.InsertOne(ctx, bson.M{"email": email, "username": username, "password": password})
	if err != nil {
		fmt.Println("Error Inserting document", err)
		return "Error Inserting document"
	}
	fmt.Println("Successful Registered user", res)
	return "Successfully Registered user: " + userDetails.Email
}
