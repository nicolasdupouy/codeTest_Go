package main

import (
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
)

func main() {
	log.Println("Initializing Hello friend API")

	router := mux.NewRouter()
	router.HandleFunc("/", ServeHTTP)

	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

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
