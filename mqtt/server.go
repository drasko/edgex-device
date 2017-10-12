package client

import (
	"net/http"

	"github.com/go-zoo/bone"
)

// HTTPServer function
func HTTPServer() http.Handler {
	mux := bone.New()

	// Status
	mux.Get("/status", http.HandlerFunc(getStatus))

	// Registration
	mux.Get("/api/v1/registration/:id", http.HandlerFunc(getRegByID))
	mux.Get("/api/v1/registration/reference/:type", http.HandlerFunc(getRegList))
	mux.Get("/api/v1/registration", http.HandlerFunc(getAllReg))
	mux.Get("/api/v1/registration/name/:name", http.HandlerFunc(getRegByName))
	mux.Post("/api/v1/registration", http.HandlerFunc(addReg))
	mux.Put("/api/v1/registration", http.HandlerFunc(updateReg))
	mux.Delete("/api/v1/registration/id/:id", http.HandlerFunc(delRegByID))
	mux.Delete("/api/v1/registration/name/:name", http.HandlerFunc(delRegByName))

	return mux
}
