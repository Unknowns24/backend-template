package handlers

import (
	"net/http"
)

// HandleBar handles requests to the /bar route
func HandleBar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bazz"))
}
