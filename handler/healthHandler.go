package handler

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type healthHandler struct{}

// HandleHealthRequest is for handling requests of kubernetes health and readiness checks
func HandleHealthRequest(router *mux.Router) {
	h := &healthHandler{}

	router.HandleFunc("/readiness", h.Health)
	router.HandleFunc("/health", h.Health)
	router.Handle("/metrics", promhttp.Handler())
}

// Health godoc
// @Summary Health is a function that stands behind the health/readiness endpoint call
// @Tags health-handler
// @Success 200 {} http.Response
// @Router /health [get]
func (*healthHandler) Health(w http.ResponseWriter, r *http.Request) {

}
