package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	port = os.Getenv("API_PORT")
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", statusHandler).Methods(http.MethodGet)

	loggedHandler := handlers.LoggingHandler(os.Stderr, router)

	server := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", "localhost", port),
		Handler: loggedHandler,
	}

	log.Printf("Listening at %v", port)

	log.Fatal(server.ListenAndServeTLS("./server.crt", "./server.key"))
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)

}
