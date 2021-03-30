package handlermethods

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"time"

	"../models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SessionCollection = models.SessionCollection

func createHash(key string) string {
	hasher := sha256.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func CreateSession(message string) (string, error) {
	result := createHash(message)
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOURI")))
	if err != nil {
		fmt.Println("Error connecting to Database", err)
		return "", errors.New("internal server error")
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
				return "", errors.New("internal server error")
			}
			fmt.Println("Session creation Successfull", res)
			return result, nil
		} else {
			fmt.Println("Error encrypting document", err)
			return "", errors.New("internal server error")
		}
	}
	fmt.Println("Session already exists")
	return result, nil
}

func GetSession(sesssionValue string) (SessionCollection, error) {
	var result SessionCollection
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOURI")))
	if err != nil {
		fmt.Println("Error connecting to Database", err)
		return result, errors.New("internal server error")
	}
	collection := client.Database("gotest").Collection("session")
	filter := bson.M{"value": sesssionValue}
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		fmt.Println("Error authencticating session", err)
		return result, errors.New("unauthorised user")
	}
	return result, nil
}

func DeleteSession(sesssionValue string) (string, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOURI")))
	if err != nil {
		fmt.Println("Error connecting to Database", err)
		return "", errors.New("internal server error")
	}
	_, err = GetSession(sesssionValue)
	if err != nil {
		fmt.Println("Error authenticating Session", err)
		return "", errors.New("internal server error")
	}
	collection := client.Database("gotest").Collection("session")
	filter := bson.M{"value": sesssionValue}
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Println("Error deleting session", err)
		return "", errors.New("internal server error")
	}
	fmt.Println("Session deletion Sucessfull", res)
	return "Session deletion Sucessfull", nil
}
