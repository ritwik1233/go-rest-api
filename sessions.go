package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

/******Session Functions START****/
type SessionCollection struct {
	ID          string    `bson:"_id"`
	Value       string    `bson:"value"`
	Email       string    `bson:"email"`
	CreatedDate time.Time `bson:"createdDate"`
}

func createSession(message, key string) (string, error) {
	result := createHash(message)
	client, err := ConnectDB()
	if err != nil {
		fmt.Println("Error connecting to Database", err)
		return "", err
	}
	collection := client.Database("gotest").Collection("session")
	// check if session exists
	var sessionData SessionCollection
	filter := bson.M{"email": message}
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	err = collection.FindOne(ctx, filter).Decode(&sessionData)
	if err != nil {
		// if session does not exists create new session
		if err.Error() == "mongo: no documents in result" {
			time := time.Now()
			res, err := collection.InsertOne(ctx, bson.M{"value": result, "email": message, "createdDate": time})
			if err != nil {
				fmt.Println("Error encrypting document", err)
				return "", err
			}
			fmt.Println("Session creation Successfull", res)
			return result, err
		} else {
			fmt.Println("Error encrypting document", err)
			return "", err
		}
	}
	fmt.Println("Session already exists")
	return result, nil
}

func getSession(sesssionValue string) (SessionCollection, error) {
	var result SessionCollection
	client, err := ConnectDB()
	if err != nil {
		fmt.Println("Error connecting to Database", err)
		return result, err
	}
	collection := client.Database("gotest").Collection("session")
	filter := bson.M{"value": sesssionValue}
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		fmt.Println("Error authencticating session", err)
		return result, err
	}
	return result, nil
}

func deleteSession(sesssionValue string) (string, error) {
	client, err := ConnectDB()
	if err != nil {
		fmt.Println("Error connecting to Database", err)
		return "", err
	}
	_, err = getSession(sesssionValue)
	if err != nil {
		fmt.Println("Error authenticating Session", err)
		return "", err
	}
	collection := client.Database("gotest").Collection("session")
	filter := bson.M{"value": sesssionValue}
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Println("Error deleting session", err)
		return "", err
	}
	fmt.Println("Session deletion Sucessfull", res)
	return "Session deletion Sucessfull", nil
}

/******Session Functions END****/
