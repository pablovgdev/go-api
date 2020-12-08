package handlers

import (
	"context"
	"go-api/data"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ProductHandler struct
type ProductHandler struct {
	log *log.Logger
}

// NewProductHandler constructor
func NewProductHandler(log *log.Logger) *ProductHandler {
	return &ProductHandler{log}
}

// GetProducts returns a slice of products
func (productHandler *ProductHandler) GetProducts(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	productHandler.log.Println("GET REQUEST")

	products := data.GetProducts()
	err := products.ToJSON(responseWriter)

	if err != nil {
		http.Error(responseWriter, "Unable to encode json", http.StatusInternalServerError)
	}
}

// AddProduct add a product to the product list
func (productHandler *ProductHandler) AddProduct(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	productHandler.log.Println("POST REQUEST")
	product := request.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(product)
}

// UpdateProduct modifies a product in the product list
func (productHandler *ProductHandler) UpdateProduct(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	productHandler.log.Println("PUT REQUEST")

	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(responseWriter, "Unable to parse product id from url", http.StatusBadRequest)
		return
	}

	product := request.Context().Value(KeyProduct{}).(*data.Product)
	err = data.UpdateProduct(id, product)

	if err != nil {
		http.Error(responseWriter, "Product not found", http.StatusBadRequest)
		return
	}
}

// KeyProduct type
type KeyProduct struct{}

// ProductValidationMiddleware parses a product from JSON and saves it in the request
func (productHandler *ProductHandler) ProductValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		if request.Method == "POST" || request.Method == "PUT" {
			product := &data.Product{}
			err := product.FromJSON(request.Body)

			if err != nil {
				http.Error(responseWriter, "Unable to decode json", http.StatusInternalServerError)
				return
			}

			ctx := context.WithValue(request.Context(), KeyProduct{}, product)
			request = request.WithContext(ctx)
		}

		next.ServeHTTP(responseWriter, request)
	})
}
