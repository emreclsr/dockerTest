package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
	})
	fmt.Println("Hello, docker")

	http.ListenAndServe(":8080", nil)
}
