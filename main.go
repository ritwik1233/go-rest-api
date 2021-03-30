package main

import (
	"net/http"

	"./handlers"
)

func main() {
	http.HandleFunc("/api/login", handlers.LoginHandler)
	http.HandleFunc("/api/register", handlers.RegisterHandler)
	http.HandleFunc("/api/logout", handlers.LogoutHandler)
	http.HandleFunc("/api/createMessage", handlers.CreateMessageHandler)
	http.HandleFunc("/api/getAllMessage", handlers.GetAllMessageHandler)
	http.HandleFunc("/api/deleteMessage", handlers.DeleteMessageHandler)
	http.HandleFunc("/api/updateMessage", handlers.UpdateMessageHandler)
	http.HandleFunc("/api/createComment", handlers.CreateCommentHandler)
	http.HandleFunc("/api/getComments", handlers.GetCommentsHandler)
	http.HandleFunc("/api/deleteComment", handlers.DeleteCommentHandler)
	http.HandleFunc("/api/updateComment", handlers.UpdateCommentHandler)
	http.ListenAndServe(":3000", nil)
}
