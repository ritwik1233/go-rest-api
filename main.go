package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeResponse(w *http.ResponseWriter, message, headerKey, headerValue string) {
	(*w).Header().Set(headerKey, headerValue)
	(*w).Write([]byte(message))
}
func loginHandler(w http.ResponseWriter, req *http.Request) {
	email := req.FormValue("email")
	password := req.FormValue("password")
	message, err := checkLoginCredentials(email, password)
	var jsonstring string
	if err != nil {
		fmt.Println("Error Logging in", err)
		jsonstring = "{\"message\":\"Login Failed\"}"
	}
	sessionValue, err := createSession(email, "session123")
	if err != nil {
		fmt.Println("Error Logging in", err)
		jsonstring = "{\"message\":\"Login Failed\"}"
	}
	jsonstring = "{\"message\":\"" + message + "\",\"auth\":\"" + sessionValue + "\"}"
	writeResponse(&w, jsonstring, "Content-Type", "application/json")
}
func registerHandler(w http.ResponseWriter, req *http.Request) {
	email := req.FormValue("email")
	password := req.FormValue("password")
	username := req.FormValue("username")
	message := registerUser(email, username, password)
	jsonstring := "{\"message\":\"" + message + "\"}"
	writeResponse(&w, jsonstring, "Content-Type", "application/json")
}
func logoutHandler(w http.ResponseWriter, req *http.Request) {
	sessionValue := req.Header.Get("Authorization")
	if len(sessionValue) == 0 {
		w.WriteHeader(401)
		jsonstring := "{\"message\":\"Unauthorized User\"}"
		writeResponse(&w, jsonstring, "Content-Type", "application/json")
		return
	}
	res, err := deleteSession(sessionValue)
	if err != nil {
		fmt.Println("Error Deleting Session", err)
		w.WriteHeader(500)
		jsonstring := "{\"message\":\" Internal Server Error \"}"
		writeResponse(&w, jsonstring, "Content-Type", "application/json")
		return
	}
	fmt.Println(res)
	jsonstring := "{\"message\":\"Logout Successfull\"}"
	writeResponse(&w, jsonstring, "Content-Type", "application/json")
}
func createMessageHandler(w http.ResponseWriter, req *http.Request) {
	sessionValue := req.Header.Get("Authorization")
	if len(sessionValue) == 0 {
		w.WriteHeader(401)
		jsonstring := "{\"message\":\"Unauthorized User\"}"
		writeResponse(&w, jsonstring, "Content-Type", "application/json")
		return
	}
	userDetails, err := getSession(sessionValue)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(401)
		jsonstring := "{\"message\":\"Unauthorized User\"}"
		writeResponse(&w, jsonstring, "Content-Type", "application/json")
		return
	}
	message := req.FormValue("message")
	res, err := CreateMessage(message, userDetails.Email)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		jsonstring := "{\"message\":\"Internal Server Error\"}"
		writeResponse(&w, jsonstring, "Content-Type", "application/json")
		return
	}
	fmt.Println(res)
	jsonstring := "{\"message\":\"Message Created Successfully\"}"
	writeResponse(&w, jsonstring, "Content-Type", "application/json")
}
func getAllMessageHandler(w http.ResponseWriter, req *http.Request) {
	sessionValue := req.Header.Get("Authorization")
	if len(sessionValue) == 0 {
		w.WriteHeader(401)
		jsonstring := "{\"message\":\"Unauthorized User\"}"
		writeResponse(&w, jsonstring, "Content-Type", "application/json")
		return
	}
	userDetails, err := getSession(sessionValue)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(401)
		jsonstring := "{\"message\":\"Unauthorized User\"}"
		writeResponse(&w, jsonstring, "Content-Type", "application/json")
		return
	}
	message, err := GetMessage(userDetails.Email)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		jsonstring := "{\"message\":\"Internal Server Error\"}"
		writeResponse(&w, jsonstring, "Content-Type", "application/json")
		return
	}
	messageData, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		jsonstring := "{\"message\":\"Internal Server Error\"}"
		writeResponse(&w, jsonstring, "Content-Type", "application/json")
		return
	}
	jsonstring := "{\"result\":" + string(messageData) + "\"}"
	writeResponse(&w, jsonstring, "Content-Type", "application/json")
}
func deleteMessageHandler(w http.ResponseWriter, req *http.Request) {
	sessionValue := req.Header.Get("Authorization")
	if len(sessionValue) == 0 {
		w.WriteHeader(401)
		jsonstring := "{\"message\":\"Unauthorized User\"}"
		writeResponse(&w, jsonstring, "Content-Type", "application/json")
		return
	}
	query := req.URL.Query().Get("messageId")
	if len(query) == 0 {
		w.WriteHeader(500)
		jsonstring := "{\"message\":\"Internal Server Error\"}"
		writeResponse(&w, jsonstring, "Content-Type", "application/json")
		return
	}
	result, err := deleteMessage(query)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		jsonstring := "{\"message\":\"Unauthorized User\"}"
		writeResponse(&w, jsonstring, "Content-Type", "application/json")
		return
	}
	fmt.Println(result)
	jsonstring := "{\"message\":\"" + result + "\"}"
	writeResponse(&w, jsonstring, "Content-Type", "application/json")

}
func main() {
	http.HandleFunc("/api/login", loginHandler)
	http.HandleFunc("/api/register", registerHandler)
	http.HandleFunc("/api/logout", logoutHandler)
	http.HandleFunc("/api/createMessage", createMessageHandler)
	http.HandleFunc("/api/getAllMessage", getAllMessageHandler)
	http.HandleFunc("/api/deleteMessage", deleteMessageHandler)
	http.ListenAndServe(":3000", nil)
}
