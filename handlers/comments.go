package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ritwik1233/go-rest-api/handlermethods"
)

func CreateCommentHandler(w http.ResponseWriter, req *http.Request) {
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
	comment := req.FormValue("comment")
	messageId := req.FormValue("messageId")
	_, err = handlermethods.CreateComment(comment, messageId, userDetails.Email)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		responsemessage := "{\"result\":\"" + err.Error() + "\"}"
		w.Write([]byte(responsemessage))
		return
	}
	responsemessage := "{\"result\":\"Comments Created Successfully\"}"
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(responsemessage))
}
func GetCommentsHandler(w http.ResponseWriter, req *http.Request) {
	sessionValue := req.Header.Get("Authorization")
	if len(sessionValue) == 0 {
		w.WriteHeader(401)
		responsemessage := "{\"result\":\"Unauthorized User\"}"
		w.Write([]byte(responsemessage))
		return
	}
	_, err := handlermethods.GetSession(sessionValue)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(401)
		responsemessage := "{\"result\":\"" + err.Error() + "\"}"
		w.Write([]byte(responsemessage))
		return
	}
	messageId := req.URL.Query().Get("messageId")
	if len(messageId) == 0 {
		w.WriteHeader(500)
		responsemessage := "{\"result\":\"Internal Server Error\"}"
		w.Write([]byte(responsemessage))
		return
	}
	message, err := handlermethods.GetComments(messageId)
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
func DeleteCommentHandler(w http.ResponseWriter, req *http.Request) {
	sessionValue := req.Header.Get("Authorization")
	if len(sessionValue) == 0 {
		w.WriteHeader(401)
		responsemessage := "{\"result\":\"Unauthorized User\"}"
		w.Write([]byte(responsemessage))
		return
	}
	queryId := req.URL.Query().Get("commentId")
	if len(queryId) == 0 {
		w.WriteHeader(500)
		responsemessage := "{\"result\":\"Internal Server Error\"}"
		w.Write([]byte(responsemessage))
		return
	}
	result, err := handlermethods.DeleteComment(queryId)
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
func UpdateCommentHandler(w http.ResponseWriter, req *http.Request) {
	sessionValue := req.Header.Get("Authorization")
	if len(sessionValue) == 0 {
		w.WriteHeader(401)
		responsemessage := "{\"result\":\"Unauthorized User\"}"
		w.Write([]byte(responsemessage))
		return
	}
	_, err := handlermethods.GetSession(sessionValue)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(401)
		responsemessage := "{\"result\":\"" + err.Error() + "\"}"
		w.Write([]byte(responsemessage))
		return
	}
	queryId := req.URL.Query().Get("commentId")
	updatedmessage := req.FormValue("comment")
	_, err = handlermethods.UpdateComment(queryId, updatedmessage)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		responsemessage := "{\"result\":\"" + err.Error() + "\"}"
		w.Write([]byte(responsemessage))
		return
	}
	responsemessage := "{\"result\":\"Comment Updated Successfully\"}"
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(responsemessage))
}
