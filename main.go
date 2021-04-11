package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/ritwik1233/go-rest-api/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ResponseMessage struct {
	Result string
}

var client *mongo.Client
var err error
var store sessions.Store

func init() {
	log.Print("Setting Environment Variable")
	env := os.Getenv("ENV")
	if env != "PROD" {
		log.Print("Loading DEV Environment")
		os.Setenv("MONGOURI", "mongodb://localhost:9001")
		os.Setenv("SESSION_KEY", "Test123")
		os.Setenv("PORT", "5000")
	} else {
		log.Print("Loading PROD Environment")
	}
	store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOURI")))
	if err != nil {
		log.Print(err.Error())
	} else {
		log.Print("Connected to Mongodb")
	}
	fs := http.FileServer(http.Dir("./client/build"))
	http.Handle("/", fs)
	http.HandleFunc("/api/register", handlers.RegisterHandler(client))
	http.HandleFunc("/api/login", handlers.LoginHandler(store, client))
	http.HandleFunc("/api/getuser", handlers.GetUser(store, client))
	http.HandleFunc("/api/logout", handlers.LogoutHandler(store))
	http.HandleFunc("/api/getAllMessage", handlers.GetAllMessageHandler(client))
	http.HandleFunc("/api/createMessage", handlers.CreateMessageHandler(store, client))
	http.HandleFunc("/api/deleteMessage", handlers.DeleteMessageHandler(store, client))
	http.HandleFunc("/api/updateMessage", handlers.UpdateMessageHandler(store, client))
	http.HandleFunc("/api/createComment", handlers.CreateCommentHandler(store, client))
	http.HandleFunc("/api/getComments", handlers.GetCommentsHandler(client))
	http.HandleFunc("/api/deleteComment", handlers.DeleteCommentHandler(store, client))
	http.HandleFunc("/api/updateComment", handlers.UpdateCommentHandler(store, client))
	PORT := ":" + os.Getenv("PORT")
	log.Print("Starting Server at PORT:", os.Getenv("PORT"))
	http.ListenAndServe(PORT, nil)
}
