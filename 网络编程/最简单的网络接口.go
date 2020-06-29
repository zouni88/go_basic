package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Header)
		fmt.Println()
		writer.Write([]byte("hello world"))
	})
	http.ListenAndServe(":8080", nil)
	
}
