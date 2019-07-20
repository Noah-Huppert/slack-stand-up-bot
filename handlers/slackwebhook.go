package handlers

import (
	"net/http"

	"github.com/Noah-Huppert/golog"
)

// SlackWebhookHandler receives Slack events
type SlackWebhookHandler struct {
	// Logger outputs debug information
	Logger golog.Logger
}

// ServeHTTP responds to HTTP requests
func (h SlackWebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
