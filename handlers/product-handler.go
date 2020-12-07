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
	} else if request.Method == http.MethodPost {
		productHandler.addProduct(responseWriter, request)
	} else {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (productHandler *ProductHandler) getProducts(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	productHandler.log.Println("GET REQUEST")

	products := data.GetProducts()
	err := products.ToJSON(responseWriter)

	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("Unable to encode json"))
	}
}

func (productHandler *ProductHandler) addProduct(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	productHandler.log.Println("POST REQUEST")

	product := &data.Product{}
	err := product.FromJSON(request.Body)

	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("Unable to decode json"))
	} else {
		data.AddProduct(product)
	}
}
