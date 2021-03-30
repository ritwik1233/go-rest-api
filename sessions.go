package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/******Session Functions START****/
type SessionCollection struct {
	ID          string    `bson:"_id"`
	Value       string    `bson:"value"`
	Email       string    `bson:"email"`
	CreatedDate time.Time `bson:"createdDate"`
}

func createHash(key string) string {
	hasher := sha256.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func CreateSession(message, key string) (string, error) {
	result := createHash(message)
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:9001"))
	if err != nil {
		fmt.Println("Error connecting to Database", err)
		return "", err
	}
	collection := client.Database("gotest").Collection("session")
	// check if session exists
	var sessionData SessionCollection
	filter := bson.M{"email": message}
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

func GetSession(sesssionValue string) (SessionCollection, error) {
	var result SessionCollection
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:9001"))
	if err != nil {
		fmt.Println("Error connecting to Database", err)
		return result, err
	}
	collection := client.Database("gotest").Collection("session")
	filter := bson.M{"value": sesssionValue}
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		fmt.Println("Error authencticating session", err)
		return result, err
	}
	return result, nil
}

func DeleteSession(sesssionValue string) (string, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:9001"))
	if err != nil {
		fmt.Println("Error connecting to Database", err)
		return "", err
	}
	_, err = GetSession(sesssionValue)
	if err != nil {
		fmt.Println("Error authenticating Session", err)
		return "", err
	}
	collection := client.Database("gotest").Collection("session")
	filter := bson.M{"value": sesssionValue}
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Println("Error deleting session", err)
		return "", err
	}
	fmt.Println("Session deletion Sucessfull", res)
	return "Session deletion Sucessfull", nil
}

/******Session Functions END****/
