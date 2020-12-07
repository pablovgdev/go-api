package handlers

import (
	"go-api/data"
	"log"
	"net/http"
)

// ProductHandler struct
type ProductHandler struct {
	log *log.Logger
}

// NewProductHandler constructor
func NewProductHandler(log *log.Logger) *ProductHandler {
	return &ProductHandler{log}
}

// ServeHTTP products handlers
func (productHandler *ProductHandler) ServeHTTP(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	if request.Method == http.MethodGet {
		productHandler.getProducts(responseWriter, request)
	}

	responseWriter.WriteHeader(http.StatusMethodNotAllowed)
}

func (productHandler *ProductHandler) getProducts(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	products := data.GetProducts()
	err := products.ToJSON(responseWriter)

	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("Unable to marshal json"))
	}
}
