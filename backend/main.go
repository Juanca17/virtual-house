package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"./connection"
	"./handler"
)

var DB = connection.Connect()

func allRight(w http.ResponseWriter, r *http.Request) {
	handler.ResponseWriter("Hello, Princess!", 200, w)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	handler.ResponseWriter("Not Found", 404, w)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", allRight).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(NotFound)
	subrouter := router.PathPrefix("/api").Subrouter()
	fmt.Println("Listening on port :1717")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "PATCH", "OPTIONS"})

	http.ListenAndServe(":1717", handlers.CORS(headersOk, originsOk, methodsOk)(router))
}
