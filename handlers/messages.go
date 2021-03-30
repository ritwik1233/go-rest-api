package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../handlermethods"
)

func CreateMessageHandler(w http.ResponseWriter, req *http.Request) {
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
	message := req.FormValue("message")
	_, err = handlermethods.CreateMessage(message, userDetails.Email)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		responsemessage := "{\"result\":\"" + err.Error() + "\"}"
		w.Write([]byte(responsemessage))
		return
	}
	responsemessage := "{\"result\":\"Message Created Successfully\"}"
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(responsemessage))
}
func GetAllMessageHandler(w http.ResponseWriter, req *http.Request) {
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
	message, err := handlermethods.GetMessage(userDetails.Email)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		responsemessage := "{\"result\":\"" + err.Error() + "\"}"
		w.Write([]byte(responsemessage))
		return
	}
	messageData, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		responsemessage := "{\"result\":\"" + err.Error() + "\"}"
		w.Write([]byte(responsemessage))
		return
	}
	responsemessage := "{\"result\":" + string(messageData) + "\"}"
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(responsemessage))
}
func DeleteMessageHandler(w http.ResponseWriter, req *http.Request) {
	sessionValue := req.Header.Get("Authorization")
	if len(sessionValue) == 0 {
		w.WriteHeader(401)
		responsemessage := "{\"result\":\"Unauthorized User\"}"
		w.Write([]byte(responsemessage))
		return
	}
	query := req.URL.Query().Get("messageId")
	if len(query) == 0 {
		w.WriteHeader(500)
		responsemessage := "{\"result\":\"Internal Server Error\"}"
		w.Write([]byte(responsemessage))
		return
	}
	result, err := handlermethods.DeleteMessage(query)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		responsemessage := "{\"result\":\"" + err.Error() + "\"}"
		w.Write([]byte(responsemessage))
		return
	}
	responsemessage := "{\"result\":\"" + result + "\"}"
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(responsemessage))
}
