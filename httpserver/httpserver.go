package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handler)
	err := http.ListenAndServe("localhost:8081", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func handler(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprint(writer, "Hello Web!")
}
