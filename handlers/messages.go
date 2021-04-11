package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/ritwik1233/go-rest-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateMessageHandler(s sessions.Store, c *mongo.Client) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Add("content-type", "application/json")
		// check if user session Exists exists
		session, err := s.Get(request, "user-session")
		if err != nil {
			log.Print(err.Error())
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		// get session value
		email := fmt.Sprintf("%s", session.Values["email"])
		if len(email) == 0 {
			log.Print("Unauthorised user")
			http.Error(response, "Unauthorised User", http.StatusForbidden)
			return
		}
		var message struct {
			Message string
		}
		json.NewDecoder(request.Body).Decode(&message)
		defer request.Body.Close()
		collection := (*c).Database("gotest").Collection("message")
		time := time.Now()
		result, err := collection.InsertOne(context.TODO(), bson.M{"message": message.Message, "createdBy": session.Values["email"], "createdDate": time})
		if err != nil {
			log.Print(err.Error())
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		var responsMessage ResponseMessage
		log.Print("Message Created successfully", result)
		myJsonString := `{"result":"Message Created successfully"}`
		json.Unmarshal([]byte(myJsonString), &responsMessage)
		json.NewEncoder(response).Encode(responsMessage)
	}
}
func GetAllMessageHandler(c *mongo.Client) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		var messages []models.MessageCollection
		response.Header().Add("content-type", "application/json")
		collection := (*c).Database("gotest").Collection("message")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			log.Print(err.Error())
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		for cursor.Next(ctx) {
			var message models.MessageCollection
			if err = cursor.Decode(&message); err != nil {
				log.Print(err.Error())
				http.Error(response, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			messages = append(messages, message)
		}
		json.NewEncoder(response).Encode(messages)
	}
}
func DeleteMessageHandler(s sessions.Store, c *mongo.Client) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Add("content-type", "application/json")
		// check if user session Exists exists
		session, err := s.Get(request, "user-session")
		if err != nil {
			log.Print(err.Error())
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		// get session value
		email := fmt.Sprintf("%s", session.Values["email"])
		if len(email) == 0 {
			log.Print("Unauthorised user")
			http.Error(response, "Unauthorised User", http.StatusForbidden)
			return
		}
		var message struct {
			Message string
		}
		json.NewDecoder(request.Body).Decode(&message)
		defer request.Body.Close()
		collection := (*c).Database("gotest").Collection("message")

		queryId := request.URL.Query().Get("messageId")
		if len(queryId) == 0 {
			log.Print(err.Error())
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		_id, err := primitive.ObjectIDFromHex(queryId)
		if err != nil {
			log.Print(err.Error())
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		var messageResult models.MessageCollection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = collection.FindOneAndDelete(ctx, bson.M{"_id": _id, "createdBy": email}).Decode(&messageResult)
		if err != nil {
			log.Print(err.Error())
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		log.Print("Message Deleted Sucessfully", messageResult)
		// delete comments if any
		collection = (*c).Database("gotest").Collection("comments")
		result, err := collection.DeleteMany(ctx, bson.M{"messageId": _id})
		if err != nil {
			log.Print(err.Error())
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		log.Print("Commeents Deleted Sucessfully", result)
		var resultMessage ResponseMessage
		myJsonString := `{"result":"Message Deleted Successfully"}`
		json.Unmarshal([]byte(myJsonString), &resultMessage)
		json.NewEncoder(response).Encode(resultMessage)
	}
}
func UpdateMessageHandler(s sessions.Store, c *mongo.Client) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Add("content-type", "application/json")
		// check if user session Exists exists
		session, err := s.Get(request, "user-session")
		if err != nil {
			log.Print(err.Error())
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		// get session value
		email := fmt.Sprintf("%s", session.Values["email"])
		if len(email) == 0 {
			log.Print("Unauthorised user")
			http.Error(response, "Unauthorised User", http.StatusForbidden)
			return
		}
		queryId := request.URL.Query().Get("messageId")
		if len(queryId) == 0 {
			log.Print(err.Error())
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		var updatedMessage struct {
			Message string
		}
		json.NewDecoder(request.Body).Decode(&updatedMessage)
		defer request.Body.Close()
		_id, err := primitive.ObjectIDFromHex(queryId)
		if err != nil {
			log.Print(err.Error())
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		collection := (*c).Database("gotest").Collection("message")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var result models.MessageCollection
		err = collection.FindOneAndUpdate(
			ctx,
			bson.M{"_id": _id, "createdBy": email},
			bson.M{
				"$set": bson.M{"message": updatedMessage.Message},
			},
		).Decode(&result)
		if err != nil {
			log.Print(err.Error())
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		var responsMessage ResponseMessage
		log.Print("Message Updated successfully", result)
		myJsonString := `{"result":"Message Updated successfully"}`
		json.Unmarshal([]byte(myJsonString), &responsMessage)
		json.NewEncoder(response).Encode(responsMessage)
	}
}
