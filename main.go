package main

import (
	"go-api/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	log := log.New(os.Stdout, "go-api", log.LstdFlags)
	productHandler := handlers.NewProductHandler(log)

	mux := mux.NewRouter()

	productsRouter := mux.PathPrefix("/products").Subrouter()
	productsRouter.Use(productHandler.ProductValidationMiddleware)
	productsRouter.HandleFunc("/", productHandler.GetProducts).Methods("GET")
	productsRouter.HandleFunc("/", productHandler.AddProduct).Methods("POST")
	productsRouter.HandleFunc("/{id:[0-9]+}", productHandler.UpdateProduct).Methods("PUT")

	http.ListenAndServe(":8000", mux)
}
