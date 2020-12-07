package handlers

import (
	"go-api/data"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"
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
	} else if request.Method == http.MethodPut {
		productHandler.updateProduct(responseWriter, request)
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
		http.Error(responseWriter, "Unable to encode json", http.StatusInternalServerError)
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
		http.Error(responseWriter, "Unable to decode json", http.StatusInternalServerError)
		return
	}

	data.AddProduct(product)
}

func (productHandler *ProductHandler) updateProduct(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	productHandler.log.Println("PUT REQUEST")

	url, err := url.Parse(request.URL.Path)

	if err != nil {
		http.Error(responseWriter, "Unable to parse products url", http.StatusBadRequest)
		return
	}

	idString := path.Base(url.Path)
	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(responseWriter, "Unable to parse product id from url", http.StatusBadRequest)
		return
	}

	product := &data.Product{}
	err = product.FromJSON(request.Body)

	if err != nil {
		http.Error(responseWriter, "Unable to decode json", http.StatusInternalServerError)
		return
	}

	err = data.UpdateProduct(id, product)

	if err != nil {
		http.Error(responseWriter, "Product not found", http.StatusBadRequest)
		return
	}
}
