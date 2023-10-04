package main

import (
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
)

func main() {
	// Create a list of backend servers.
	servers := []string{
		// By not specifying an IP address before the colon, the server will listen
		// on every IP address associated with your computer
		":8081",
		":8082",
		":8083",
	}

	// Create a reverse proxy.
	proxy := httputil.NewSingleHostReverseProxy(nil)

	// Create a handler that forwards requests to the reverse proxy.
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Select a backend service
		targetBackendService := servers[rand.Intn(3)]

		// Set the target backend service.
		proxy.Director = func(r *http.Request) {
			// the Director is just function that modifies the original incoming request
			r.URL.Scheme = "http"
			r.URL.Host = targetBackendService
		}

		// Serve the request
		proxy.ServeHTTP(w, r)
	})

	// Start the load balancer.
	log.Fatal(http.ListenAndServe(":8080", handler))
}
