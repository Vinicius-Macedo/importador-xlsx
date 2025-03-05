package handlers

import (
	"log"
	"net/http"
)

// GetUser retrieves the authenticated user's information.
// @Summary Get authenticated user
// @Description Retrieve the authenticated user's information using cookie HttpOnly JWT.
// @Tags user
// @Produce json
// @Success 200 {object} GetUserResponse
// @Failure 401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /user [get]
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := h.JWT.ExtractUserID(r)

	if err != nil {
		log.Println("error extracting user id from token", err)
		h.errorJSON(w, http.StatusUnauthorized, "invalid token")
		return
	}

	user, err := h.services.GetUserByID(userID)

	if err != nil {
		h.errorJSON(w, http.StatusInternalServerError, "Failed to get user info")
		return
	}

	h.writeJSON(w, http.StatusOK, GetUserResponse{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
	})
}
