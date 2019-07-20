package handlers

import (
	"fmt"
	"net/http"

	"github.com/Noah-Huppert/golog"
)

// HealthHandler returns status code 200 if API is OK
type HealthCheckHandler struct {
	// Logger outputs debug information
	Logger golog.Logger
}

// ServeHTTP responds to HTTP requests
func (h HealthCheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "{\"ok\": true}")
	if err != nil {
		h.Logger.Errorf("failed to write HTTP response: %s", err.Error())
	}
}
