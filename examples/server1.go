package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// Telling the server to call rootHandler when the '/' URL is hit
	http.HandleFunc("/", rootHandler)

	// When nil is set for the 'http.Handler' parameter, the server will use the default multiplexer
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// w http.ResponseWriter, r *http.Request signature used by http handler functions
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got request from server 1!")
	io.WriteString(w, "Hello from server 1")
}
