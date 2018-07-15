package handlers

import (
	"fmt"
	"net/http"
	"os"
)

type HelloResponse struct {
	Message string
	Version string
}

type helloWorldHandler struct {
	Message string
	Version string
}

func (h *helloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := HelloResponse{
		Message: h.Message,
		Version: h.Version,
	}
	// Get hostname
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "%s, I am running on %q and my version is %q.", response.Message, hostname, response.Version)
}

func HelloWorldHandler(message string, version string) http.Handler {
	if message == "" {
		message = "Hello OSCON!"
	}

	return &helloWorldHandler{
		Message: message,
		Version: version,
	}
}
