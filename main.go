package main

import (
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "GET method requested"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "POST method requested"}`))
	case "PUT":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "PUT method requested"}`))
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "DELETE method requested"}`))
	case "PATCH":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "PATCH method requested"}`))
	case "OPTIONS":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "OPTIONS method requested"}`))
	case "HEAD":
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Can't find method requested"}`))
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
