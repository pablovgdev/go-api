package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World!")
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Bad Request"))
			log.Fatal(err.Error())
			return
		}
		log.Print(string(data))
		rw.Write(data)
	})
	http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye World!")
	})
	http.ListenAndServe(":8000", nil)
}
