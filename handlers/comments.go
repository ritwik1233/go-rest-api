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

func CreateCommentHandler(s sessions.Store, c *mongo.Client) http.HandlerFunc {
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
		var comment struct {
			Comment   string
			MessageId string
		}
		json.NewDecoder(request.Body).Decode(&comment)
		defer request.Body.Close()
		collection := (*c).Database("gotest").Collection("message")
		var messages models.MessageCollection
		_id, err := primitive.ObjectIDFromHex(comment.MessageId)
		if err != nil {
			log.Print(err.Error())
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = collection.FindOne(ctx, bson.M{"_id": _id}).Decode(&messages)
		if err != nil {
			log.Print(err.Error())
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		time := time.Now()
		collection = (*c).Database("gotest").Collection("comments")
		result, err := collection.InsertOne(ctx, bson.M{"comment": comment.Comment, "messageId": _id, "createdBy": email, "createdDate": time})
		if err != nil {
			log.Print(err.Error())
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		var responsMessage ResponseMessage
		log.Print("Comment Created successfully", result)
		myJsonString := `{"result":"Comment Created successfully"}`
		json.Unmarshal([]byte(myJsonString), &responsMessage)
		json.NewEncoder(response).Encode(responsMessage)
	}
}
func GetCommentsHandler(c *mongo.Client) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Add("content-type", "application/json")
		var comments []models.CommentsCollection
		queryId := request.URL.Query().Get("messageId")
		if len(queryId) == 0 {
			log.Print("Empty query")
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		_id, err := primitive.ObjectIDFromHex(queryId)
		if err != nil {
			log.Print(err.Error())
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		collection := (*c).Database("gotest").Collection("comments")
		cursor, err := collection.Find(ctx, bson.M{"messageId": _id})
		if err != nil {
			log.Print(err.Error())
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		for cursor.Next(ctx) {
			var comment models.CommentsCollection
			if err = cursor.Decode(&comment); err != nil {
				log.Print(err.Error())
				http.Error(response, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			comments = append(comments, comment)
		}
		json.NewEncoder(response).Encode(comments)
	}
}
func DeleteCommentHandler(s sessions.Store, c *mongo.Client) http.HandlerFunc {
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
		queryId := request.URL.Query().Get("commentId")
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
		collection := (*c).Database("gotest").Collection("comments")
		var commentResult models.CommentsCollection
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = collection.FindOneAndDelete(ctx, bson.M{"_id": _id, "createdBy": email}).Decode(&commentResult)
		if err != nil {
			log.Print(err.Error())
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		log.Print("Comment Deleted Sucessfully", commentResult)
		var resultMessage ResponseMessage
		myJsonString := `{"result":"Message Deleted Successfully"}`
		json.Unmarshal([]byte(myJsonString), &resultMessage)
		json.NewEncoder(response).Encode(resultMessage)
	}
}
func UpdateCommentHandler(s sessions.Store, c *mongo.Client) http.HandlerFunc {
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
		queryId := request.URL.Query().Get("commentId")
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

		var updatedComment struct {
			Comment string
		}
		json.NewDecoder(request.Body).Decode(&updatedComment)
		defer request.Body.Close()
		collection := (*c).Database("gotest").Collection("comments")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var result models.CommentsCollection
		err = collection.FindOneAndUpdate(
			ctx,
			bson.M{"_id": _id, "createdBy": email},
			bson.M{
				"$set": bson.M{"comment": updatedComment.Comment},
			},
		).Decode(&result)
		if err != nil {
			log.Print(err.Error())
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		var responsMessage ResponseMessage
		log.Print("Comment Updated successfully", result)
		myJsonString := `{"result":"Comment Updated successfully"}`
		json.Unmarshal([]byte(myJsonString), &responsMessage)
		json.NewEncoder(response).Encode(responsMessage)
	}
}
