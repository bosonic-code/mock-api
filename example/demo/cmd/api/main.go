package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/nicklaw5/helix"
)

var (
	port           = os.Getenv("API_PORT")
	twitchClientID = os.Getenv("TWITCH_CLIENT_ID")
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/users/{id}/is-streaming", statusHandler).Methods(http.MethodGet)

	loggedHandler := handlers.LoggingHandler(os.Stderr, router)

	server := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", "localhost", port),
		Handler: loggedHandler,
	}

	log.Printf("Listening at %v", port)

	log.Fatal(server.ListenAndServe())
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	client, err := helix.NewClient(&helix.Options{
		ClientID: twitchClientID,
	})

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	userID := mux.Vars(r)["id"]

	resp, err := client.GetStreams(&helix.StreamsParams{UserIDs: []string{userID}})

	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	if resp.StatusCode != 200 {
		w.WriteHeader(500)
		log.Println(resp.StatusCode)
		log.Println(resp.ErrorMessage)
		log.Println(resp.ErrorStatus)
		return
	}

	w.WriteHeader(200)
	if len(resp.Data.Streams) == 0 {
		w.Write([]byte("{\"isStreaming\" : false}"))
	} else {
		w.Write([]byte("{\"isStreaming\" : true}"))
	}

}
