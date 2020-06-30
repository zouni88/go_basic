package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		writer.Write([]byte("hello 111world"))
	})
	log.Fatal(http.ListenAndServe(":8080", router))
}
