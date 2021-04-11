package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/ritwik1233/go-rest-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type ResponseMessage struct {
	Result string
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func RegisterHandler(c *mongo.Client) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Add("content-type", "application/json")
		var user models.UserCollection
		json.NewDecoder(request.Body).Decode(&user)
		defer request.Body.Close()
		password, err := HashPassword(user.Password)
		if err != nil {
			log.Print(err.Error())
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		user.Password = password
		collection := (*c).Database("gotest").Collection("users")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var result models.UserCollection
		// check if user exists
		filter := bson.M{"email": user.Email}
		err = collection.FindOne(ctx, filter).Decode(&result)
		if err != nil {
			noDocumentMessage := "mongo: no documents in result"
			if err.Error() != noDocumentMessage {
				log.Print(err.Error())
				http.Error(response, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			result, err := collection.InsertOne(ctx, user)
			if err != nil {
				log.Print(err.Error())
				http.Error(response, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			var resultMessage ResponseMessage
			log.Print("User Registered Successfully", result.InsertedID)
			myJsonString := `{"result":"User Registered Successfully"}`
			json.Unmarshal([]byte(myJsonString), &resultMessage)
			json.NewEncoder(response).Encode(resultMessage)
		}
		if len(result.Email) > 0 {
			var errMessage ResponseMessage
			log.Print("User Already Exists")
			myJsonString := `{"result":"User Already Exists"}`
			json.Unmarshal([]byte(myJsonString), &errMessage)
			json.NewEncoder(response).Encode(errMessage)
		}
	}
}

func LoginHandler(s sessions.Store, c *mongo.Client) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Add("content-type", "application/json")
		var user struct {
			Email    string
			Password string
		}
		json.NewDecoder(request.Body).Decode(&user)
		defer request.Body.Close()
		password, err := HashPassword(user.Password)
		user.Password = password
		if err != nil {
			log.Print(err.Error())
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		collection := (*c).Database("gotest").Collection("users")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var result models.UserCollection
		// check if user exists
		filter := bson.M{"email": user.Email}
		err = collection.FindOne(ctx, filter).Decode(&result)
		if err != nil {
			log.Print(err.Error())
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		session, _ := s.Get(request, "user-session")
		session.Values["email"] = user.Email
		err = session.Save(request, response)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}
		var resultMessage ResponseMessage
		log.Print("Login Successfull", session.Values["email"])
		myJsonString := `{"result":"Login Successfull"}`
		json.Unmarshal([]byte(myJsonString), &resultMessage)
		json.NewEncoder(response).Encode(resultMessage)
	}
}

func GetUser(s sessions.Store, c *mongo.Client) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Add("content-type", "application/json")
		collection := (*c).Database("gotest").Collection("users")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var result models.UserCollection
		// check if user session Exists exists
		session, err := s.Get(request, "user-session")
		if err != nil {
			log.Print(err.Error())
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		filter := bson.M{"email": session.Values["email"]}
		err = collection.FindOne(ctx, filter).Decode(&result)
		if err != nil {
			noDocumentMessage := "mongo: no documents in result"
			if err.Error() != noDocumentMessage {
				log.Print(err.Error())
				http.Error(response, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			var errMessage ResponseMessage
			log.Print("No User Exists")
			myJsonString := `{"result":"No User Exists"}`
			json.Unmarshal([]byte(myJsonString), &errMessage)
			json.NewEncoder(response).Encode(errMessage)
			return
		}
		result.Password = ""
		log.Println("Session exists")
		json.NewEncoder(response).Encode(result)
	}
}

func LogoutHandler(s sessions.Store) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Add("content-type", "application/json")
		// check if user session Exists exists
		session, err := s.Get(request, "user-session")
		if err != nil {
			log.Print(err.Error())
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		session.Values["email"] = ""
		err = session.Save(request, response)
		if err != nil {
			log.Print(err.Error())
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("Session removed", session.Values["email"])
		var resultMessage ResponseMessage
		log.Print("Logout Successfull", session.Values["email"])
		myJsonString := `{"result":"Logout Successfull"}`
		json.Unmarshal([]byte(myJsonString), &resultMessage)
		json.NewEncoder(response).Encode(resultMessage)
	}
}
