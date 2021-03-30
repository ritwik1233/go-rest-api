package handlers

import (
	"fmt"
	"net/http"

	"../handlermethods"
)

func LoginHandler(w http.ResponseWriter, req *http.Request) {
	email := req.FormValue("email")
	password := req.FormValue("password")
	message, err := handlermethods.CheckLoginCredentials(email, password)
	var responsemessage string
	if err != nil {
		w.WriteHeader(404)
		fmt.Println("Error Logging in", err)
		responsemessage = "{\"result\":\"" + err.Error() + "\"}"
		w.Write([]byte(responsemessage))
		return
	}
	sessionValue, err := handlermethods.CreateSession(email, "session123")
	if err != nil {
		w.WriteHeader(404)
		fmt.Println("Error Logging in", err)
		responsemessage = "{\"result\":\"" + err.Error() + "\"}"
		w.Write([]byte(responsemessage))
		return
	}
	responsemessage = "{\"message\":\"" + message + "\",\"auth\":\"" + sessionValue + "\"}"
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(responsemessage))
}
func RegisterHandler(w http.ResponseWriter, req *http.Request) {
	email := req.FormValue("email")
	password := req.FormValue("password")
	username := req.FormValue("username")
	message, err := handlermethods.RegisterUser(email, username, password)
	if err != nil {
		fmt.Println("Error Registering User", err)
		w.WriteHeader(500)
		responsemessage := "{\"result\":\"" + err.Error() + "\"}"
		w.Write([]byte(responsemessage))
		return
	}
	responsemessage := "{\"result\":\"" + message + "\"}"
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(responsemessage))
}
func LogoutHandler(w http.ResponseWriter, req *http.Request) {
	sessionValue := req.Header.Get("Authorization")
	if len(sessionValue) == 0 {
		w.WriteHeader(401)
		responsemessage := "{\"result\":\"Unauthorized User\"}"
		w.Write([]byte(responsemessage))
		return
	}
	_, err := handlermethods.DeleteSession(sessionValue)
	if err != nil {
		fmt.Println("Error Deleting Session", err)
		w.WriteHeader(500)
		responsemessage := "{\"result\":\"" + err.Error() + "\"}"
		w.Write([]byte(responsemessage))
		return
	}
	responsemessage := "{\"result\":\"Logout Successfull\"}"
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(responsemessage))
}
