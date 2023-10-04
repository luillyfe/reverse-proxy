package main

import "luillyfe.com/reverse-proxy/http"

func main() {
	// Collecting a list of backend servers.
	endpoints := []string{
		// By not specifying an IP address before the colon, the server will listen
		// on every IP address associated with your computer
		":8081",
		":8082",
		":8083",
	}

	// Create a Load Balancer
	loadBalancer := &http.ReverseProxy{
		BackendService: http.BackendService{Backend: endpoints},
	}

	// Run the Load Balancer
	loadBalancer.Run()
}
