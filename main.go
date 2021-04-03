package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ritwik1233/go-rest-api/dev"
	"github.com/ritwik1233/go-rest-api/handlers"
)

func defaultHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("<h1>Default Page</h1>"))
}

func init() {
	fmt.Println("Setting Environment Variable")
	env := os.Getenv("ENV")
	if env != "PROD" {
		fmt.Println("Loading Dev Environment")
		var devkeys dev.Keys
		devkeys.Initialize()
		os.Setenv("MONGOURI", devkeys.MongoURI)
		os.Setenv("PORT", devkeys.PORT)
		return
	}
	fmt.Println("Loading PROD Environment")
}
func main() {
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
	http.HandleFunc("/", defaultHandler)
	PORT := ":" + os.Getenv("PORT")
	fmt.Println("Starting Server at PORT:", os.Getenv("PORT"))
	http.ListenAndServe(PORT, nil)
}
