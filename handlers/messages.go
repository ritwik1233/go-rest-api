package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../handlermethods"
)

func writeResponse(w *http.ResponseWriter, message, headerKey, headerValue string) {
	(*w).Header().Set(headerKey, headerValue)
	(*w).Write([]byte(message))
}
func CreateMessageHandler(w http.ResponseWriter, req *http.Request) {
	sessionValue := req.Header.Get("Authorization")
	if len(sessionValue) == 0 {
		w.WriteHeader(401)
		jsonstring := "{\"message\":\"Unauthorized User\"}"
		writeResponse(&w, jsonstring, "Content-Type", "application/json")
		return
	}
	userDetails, err := handlermethods.GetSession(sessionValue)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(401)
		jsonstring := "{\"message\":\"Unauthorized User\"}"
		writeResponse(&w, jsonstring, "Content-Type", "application/json")
		return
	}
	message := req.FormValue("message")
	res, err := handlermethods.CreateMessage(message, userDetails.Email)
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
func GetAllMessageHandler(w http.ResponseWriter, req *http.Request) {
	sessionValue := req.Header.Get("Authorization")
	if len(sessionValue) == 0 {
		w.WriteHeader(401)
		jsonstring := "{\"message\":\"Unauthorized User\"}"
		writeResponse(&w, jsonstring, "Content-Type", "application/json")
		return
	}
	userDetails, err := handlermethods.GetSession(sessionValue)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(401)
		jsonstring := "{\"message\":\"Unauthorized User\"}"
		writeResponse(&w, jsonstring, "Content-Type", "application/json")
		return
	}
	message, err := handlermethods.GetMessage(userDetails.Email)
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
func DeleteMessageHandler(w http.ResponseWriter, req *http.Request) {
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
	result, err := handlermethods.DeleteMessage(query)
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
