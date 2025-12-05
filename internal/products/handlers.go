package products

import (
	"ecom/internal/json"
	"log"
	"net/http"
)

type handler struct {
	service Service
}

// Dependency injection - here service is a dependency for the Handler
// This is like a factory function to instantiate new handlers

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {

	err := h.service.ListProducts(r.Context())

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	products := struct {
		Products []string `json:"products"`
	}{}

	json.Write(w, http.StatusOK, products)
}
