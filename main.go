package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ritwik1233/go-rest-api/handlers"
)

func init() {
	fmt.Println("Setting Environment Variable")
	env := os.Getenv("ENV")
	if env != "PROD" {
		fmt.Println("Loading DEV Environment")
		os.Setenv("MONGOURI", "mongodb://localhost:9001")
		os.Setenv("PORT", "5000")
		return
	} else {
		fmt.Println("Loading PROD Environment")
	}
}

func main() {
	fs := http.FileServer(http.Dir("./client/build"))
	http.Handle("/", fs)
	http.HandleFunc("/api/login", handlers.LoginHandler)
	http.HandleFunc("/api/register", handlers.RegisterHandler)
	http.HandleFunc("/api/logout", handlers.LogoutHandler)
	http.HandleFunc("/api/getCurrentUser", handlers.GetCurrenUser)
	http.HandleFunc("/api/createMessage", handlers.CreateMessageHandler)
	http.HandleFunc("/api/getAllMessage", handlers.GetAllMessageHandler)
	http.HandleFunc("/api/deleteMessage", handlers.DeleteMessageHandler)
	http.HandleFunc("/api/updateMessage", handlers.UpdateMessageHandler)
	http.HandleFunc("/api/createComment", handlers.CreateCommentHandler)
	http.HandleFunc("/api/getComments", handlers.GetCommentsHandler)
	http.HandleFunc("/api/deleteComment", handlers.DeleteCommentHandler)
	http.HandleFunc("/api/updateComment", handlers.UpdateCommentHandler)
	PORT := ":" + os.Getenv("PORT")
	fmt.Println("Starting Server at PORT:", os.Getenv("PORT"))
	http.ListenAndServe(PORT, nil)
}
