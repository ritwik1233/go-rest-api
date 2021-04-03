package handlers

import (
	"fmt"
	"net/http"

	"github.com/ritwik1233/go-rest-api/handlermethods"
)

func LoginHandler(w http.ResponseWriter, req *http.Request) {
	var responsemessage string
	email := req.FormValue("email")
	password := req.FormValue("password")
	message, err := handlermethods.CheckLoginCredentials(email, password)
	if err != nil {
		w.WriteHeader(404)
		fmt.Println("Error Logging in", err)
		responsemessage = "{\"result\":\"" + err.Error() + "\"}"
		w.Write([]byte(responsemessage))
		return
	}
	sessionValue, err := handlermethods.CreateSession(email)
	if err != nil {
		w.WriteHeader(404)
		fmt.Println("Error Logging in", err)
		responsemessage = "{\"result\":\"" + err.Error() + "\"}"
		w.Write([]byte(responsemessage))
		return
	}
	responsemessage = "{\"result\":\"" + message + "\",\"auth\":\"" + sessionValue + "\"}"
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(responsemessage))
}
func RegisterHandler(w http.ResponseWriter, req *http.Request) {
	email := req.FormValue("email")
	password := req.FormValue("password")
	message, err := handlermethods.RegisterUser(email, password)
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

func GetCurrenUser(w http.ResponseWriter, req *http.Request) {
	sessionValue := req.Header.Get("Authorization")
	if len(sessionValue) == 0 {
		w.WriteHeader(401)
		responsemessage := "{\"result\":\"Unauthorized User\"}"
		w.Write([]byte(responsemessage))
		return
	}
	userDetails, err := handlermethods.GetSession(sessionValue)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(401)
		responsemessage := "{\"result\":\"" + err.Error() + "\"}"
		w.Write([]byte(responsemessage))
		return
	}
	responsemessage := "{\"result\":\"" + userDetails.Email + "\"}"
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(responsemessage))
}
