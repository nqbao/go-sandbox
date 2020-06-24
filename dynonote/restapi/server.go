package restapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/nqbao/learn-go/dynonote/model"
	"github.com/nqbao/learn-go/dynonote/service"
)

const address = ":5000"

type ListNoteResponses struct {
	Items             []*model.Note
	LastEvaluationKey string
}

func getUser(r *http.Request) string {
	return "bao" // XXX
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
	writer := json.NewEncoder(w)

	vars := mux.Vars(r)
	ts, _ := strconv.Atoi(vars["timestamp"])
	err := service.DeleteNote(getUser(r), ts)

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

func listUserNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	limitArg := r.URL.Query()["limit"]
	limit := 20

	if len(limitArg) > 1 {
		limit, _ = strconv.Atoi(limitArg[0])
	}

	notes, err := service.GetUserNote(getUser(r), limit, vars["next"])

	writer := json.NewEncoder(w)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writer.Encode(map[string]string{
			"status":  "error",
			"message": fmt.Sprintf("%v", err),
		})
	} else {
		repr := ListNoteResponses{
			Items: notes,
		}
		writer.Encode(repr)
	}
}

func StartServer() {
	router := mux.NewRouter().StrictSlash(true)

	router.Path("/api/note").
		Methods("POST").
		HandlerFunc(postNote)

	router.Path("/api/note").
		Methods("PATCH").
		HandlerFunc(updateNote)

	router.Path("/api/note/{timestamp}").
		Methods("DELETE").
		HandlerFunc(deleteNote)

	router.Path("/api/notes").
		Queries("next", "{next}").
		HandlerFunc(listUserNote)

	router.HandleFunc("/api/notes", listUserNote)

	router.Path("/api/tokensignin").Methods("POST").HandlerFunc(tokenSignIn)

	log.Printf("Listening on address %s", address)

	corsRoute := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"content-type", "authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "HEAD", "PATCH", "DELETE"}),
	)(authHandler(router))
	loggingRouter := handlers.LoggingHandler(os.Stdout, corsRoute)

	log.Fatal(http.ListenAndServe(address, loggingRouter))
}
