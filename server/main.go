package main

import (
	"fmt"
	"log"
	"net/http"

	h2c "github.com/http2-h2c/common"
)

func main() {
	// exmaple for HTTP/2 h2c
	// https://github.com/golang/go/issues/14141
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello message")
	})

	fmt.Println("Server started...")
	log.Fatal(http.ListenAndServe(":8080", &h2c.Server{}))
	// regular golang implementation
	// support only HTTP/2 TLS
	// log.Fatal(http.ListenAndServe(":8080", nil))
}
