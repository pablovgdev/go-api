package main

import (
	"go-api/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	log := log.New(os.Stdout, "go-api", log.LstdFlags)
	indexHandler := handlers.NewIndexHandler(log)
	productHandler := handlers.NewProductHandler(log)

	mux := http.NewServeMux()
	mux.Handle("/", indexHandler)
	mux.Handle("/products/", productHandler)

	http.ListenAndServe(":8000", mux)
}
