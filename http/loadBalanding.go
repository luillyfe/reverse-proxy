package http

import (
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
)

type BackendService struct {
	Servers []string
}

type ReverseProxy struct {
	BackendService BackendService
}

func (rp *ReverseProxy) Run() {
	// Create a reverse proxy.
	proxy := httputil.NewSingleHostReverseProxy(nil)

	// Create a handler that forwards requests to the reverse proxy.
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Select a backend service
		targetBackendService := rp.selectService()

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

// Load balancing algorithm
func (rp *ReverseProxy) selectService() string {
	return rp.BackendService.Servers[rand.Intn(3)]
}
