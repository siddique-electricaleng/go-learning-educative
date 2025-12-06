package products

import (
	"ecom/internal/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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

	products, err := h.service.ListProducts(r.Context())

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, products)
}

func (h *handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	// 1. Extract the id from the URL
	idParam := chi.URLParam(r, "id")

	// 2. Convert the id to an integer
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Id", http.StatusBadRequest)
	}

	// 3. Call service layer
	product, err := h.service.GetProductByID(r.Context(), id)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Write(w, http.StatusOK, product)
}
