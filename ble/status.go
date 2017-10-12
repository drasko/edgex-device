package client

import (
	"io"
	"net/http"
)

func getStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	str := `{"running": true}`
	io.WriteString(w, str)
}
