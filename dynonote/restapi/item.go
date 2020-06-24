package restapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nqbao/learn-go/dynonote/model"
	"github.com/nqbao/learn-go/dynonote/service"
)

type PostItemRequest struct {
	Item *model.Note
}

func postNote(w http.ResponseWriter, r *http.Request) {
	nm := service.NewNoteManager(getAwsCredentials(r))

	req := &PostItemRequest{}
	writer := json.NewEncoder(w)
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writer.Encode(map[string]string{
			"status":  "error",
			"message": fmt.Sprintf("%v", err),
		})
	} else {
		note := req.Item
		note.UserKey = getUser(r)
		note.UserName = getUserName(r)

		err := nm.CreateNote(note)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			writer.Encode(map[string]string{
				"status":  "error",
				"message": fmt.Sprintf("%v", err),
			})
		} else {
			writer.Encode(note)
		}
	}
}

func updateNote(w http.ResponseWriter, r *http.Request) {
	nm := service.NewNoteManager(getAwsCredentials(r))

	req := &PostItemRequest{}
	writer := json.NewEncoder(w)
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writer.Encode(map[string]string{
			"status":  "error",
			"message": fmt.Sprintf("%v", err),
		})
	} else {
		note := req.Item
		note.UserKey = getUser(r)
		note.UserName = getUserName(r)

		err := nm.UpdateNote(note)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			writer.Encode(map[string]string{
				"status":  "error",
				"message": fmt.Sprintf("%v", err),
			})
		} else {
			writer.Encode(note)
		}
	}
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
	nm := service.NewNoteManager(getAwsCredentials(r))
	writer := json.NewEncoder(w)

	vars := mux.Vars(r)
	ts, _ := strconv.Atoi(vars["timestamp"])
	err := nm.DeleteNote(getUser(r), ts)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writer.Encode(map[string]string{
			"status":  "error",
			"message": fmt.Sprintf("%v", err),
		})
	} else {
		w.Write([]byte("{\"success\": true}"))
	}
}
