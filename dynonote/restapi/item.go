package restapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nqbao/learn-go/dynonote/model"
	"github.com/nqbao/learn-go/dynonote/service"
)

type PostItemRequest struct {
	Item *model.Note
}

func postNote(w http.ResponseWriter, r *http.Request) {
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
		note.UserName = getUser(r)

		err := service.CreateNote(note)

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

		err := service.UpdateNote(note)

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
