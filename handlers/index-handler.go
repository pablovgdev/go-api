package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
)

// IndexHandler struct
type IndexHandler struct {
	log *log.Logger
}

// NewIndexHandler constructor
func NewIndexHandler(log *log.Logger) *IndexHandler {
	return &IndexHandler{log}
}

// ServeHTTP all requests
func (indexHandler *IndexHandler) ServeHTTP(
	responseWriter http.ResponseWriter,
	request *http.Request,
) {
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte("Bad Request"))
		indexHandler.log.Fatal(err.Error())
	} else {
		indexHandler.log.Println("Server listening")
		indexHandler.log.Printf("Data: %s\n", data)
		responseWriter.Write(data)
	}
}
