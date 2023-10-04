package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootHandler)

	log.Fatal(http.ListenAndServe(":8083", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got request from server 3!")
	io.WriteString(w, "Hello from server 3")
}
