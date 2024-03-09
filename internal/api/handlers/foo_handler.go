package handlers

import (
	"net/http"
)

// HandleFoo handles requests to the /foo route
func HandleFoo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bar"))
}
