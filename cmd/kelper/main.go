package main

import (
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	router := NewRouter()
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	log.Fatal(http.ListenAndServe(":8000", loggedRouter))
}
