package handlers

import (
	"log"
	"net/http"
)

// GetAllCustomers is a handler that returns all customers
// @Summary Get all customers
// @Description Get all customers, route protected by JWT
// @Tags Get xlsx data
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errorResponse
// @Router /customers [get]
func (h *Handler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	rows, err := h.services.GetAllCustomers()
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	h.writeJSON(w, 200, rows)
}

// GetAllResources is a handler that returns all resources
// @Summary Get all resources
// @Description Get all resources, route protected by JWT
// @Tags Get xlsx data
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errorResponse
// @Router /resources [get]
func (h *Handler) GetAllResources(w http.ResponseWriter, r *http.Request) {
	rows, err := h.services.GetAllResources()
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	h.writeJSON(w, 200, rows)
}

// Get all categories
// @Summary Get all categories
// @Description Get all categories, route protected by JWT
// @Tags Get xlsx data
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} errorResponse
// @Router /categories [get]
func (h *Handler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	rows, err := h.services.GetAllCategories()
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	h.writeJSON(w, 200, rows)
}
