package main

import (
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
)

func main() {
	log.Println("Initializing Hello friend API")

	router := mux.NewRouter()
	router.HandleFunc("/", ServeHTTP)

	server := http.Server{
		Addr:    net.JoinHostPort("", "8080"),
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// Write a JSON message into the response body
	message := `{"message": "Hello friend !"}`
	_, err := w.Write([]byte(message))

	if err != nil {
		log.Fatal(err)
	}
}
