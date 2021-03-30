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
	http.ListenAndServe(":3000", nil)
}
