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
	var jsonstring string
	if err != nil {
		fmt.Println("Error Logging in", err)
		jsonstring = "{\"message\":\"Login Failed\"}"
	}
	sessionValue, err := handlermethods.CreateSession(email, "session123")
	if err != nil {
		fmt.Println("Error Logging in", err)
		jsonstring = "{\"message\":\"Login Failed\"}"
	}
	jsonstring = "{\"message\":\"" + message + "\",\"auth\":\"" + sessionValue + "\"}"
	writeResponse(&w, jsonstring, "Content-Type", "application/json")
}
func RegisterHandler(w http.ResponseWriter, req *http.Request) {
	email := req.FormValue("email")
	password := req.FormValue("password")
	username := req.FormValue("username")
	message, err := handlermethods.RegisterUser(email, username, password)
	if err != nil {
		fmt.Println("Error Registering User", err)
		w.WriteHeader(500)
		jsonstring := "{\"message\":\" Internal Server Error \"}"
		writeResponse(&w, jsonstring, "Content-Type", "application/json")
		return
	}
	jsonstring := "{\"message\":\"" + message + "\"}"
	writeResponse(&w, jsonstring, "Content-Type", "application/json")
}
func LogoutHandler(w http.ResponseWriter, req *http.Request) {
	sessionValue := req.Header.Get("Authorization")
	if len(sessionValue) == 0 {
		w.WriteHeader(401)
		jsonstring := "{\"message\":\"Unauthorized User\"}"
		writeResponse(&w, jsonstring, "Content-Type", "application/json")
		return
	}
	res, err := handlermethods.DeleteSession(sessionValue)
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
