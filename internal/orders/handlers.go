package orders

import (
	"ecom/internal/json"
	"log"
	"net/http"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) PlaceOrder(w http.ResponseWriter, r *http.Request) {

	// Read and Pass/Marshal the Payload: json body into a Go struct
	var tempOrder createOrderParams

	if err := json.Read(r, &tempOrder); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call the service: to post this struct into the sql
	createdOrder, err := h.service.PlaceOrder(r.Context(), tempOrder)
	if err != nil {
		log.Println(err)
		if err == ErrProductNotFound || err == ErrProductNoStock {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return to the user a response body with a status code
	json.Write(w, http.StatusCreated, createdOrder)
}
