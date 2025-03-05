package handlers

import (
	"api/cmd/internal/auth"
	"api/cmd/internal/helpers"
	"api/cmd/internal/postgresrepo"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type createUserResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// CreateUser creates a new user with the provided data.
// @Summary Create a new user
// @Description Create a new user with the provided data. The password must be at least 8 characters long, contain an uppercase letter, a lowercase letter, and a symbol.
// @Tags auth
// @Accept json
// @Produce json
// @Param user body postgresrepo.CreateUserParams true "User data"
// @Success 201 {object} postgresrepo.CreateUserParams
// @Failure 400 {object} errorResponse
// @Failure 409 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /register [post]
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userParams postgresrepo.CreateUserParams
	err := json.NewDecoder(r.Body).Decode(&userParams)

	if err != nil {
		h.errorJSON(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	fields := map[string]string{
		"name":     userParams.Name,
		"username": userParams.Username,
		"email":    userParams.Email,
		"password": userParams.Password,
	}

	message, err := helpers.ValidateFields(fields)

	if err != nil {
		log.Println("error validating fields", message)
		h.errorJSON(w, http.StatusBadRequest, message)
		return
	}

	exists, err := h.services.CheckIfUserExistsByEmail(userParams.Email)

	if err != nil {
		h.errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if exists {
		log.Println("error email already exists")
		h.errorJSON(w, http.StatusConflict, "o email já está em uso")
		return
	}

	exists, err = h.services.CheckIfUserExistsByUsername(userParams.Username)

	if err != nil {
		h.errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if exists {
		log.Println("error username already exists")
		h.errorJSON(w, http.StatusConflict, "o nome de usuário já está em uso")
		return
	}

	hashedPassword, err := h.hashPassword(userParams.Password)

	if err != nil {
		h.errorJSON(w, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	user := postgresrepo.CreateUserParams{
		Name:     userParams.Name,
		Username: userParams.Username,
		Email:    userParams.Email,
		Password: hashedPassword,
	}

	err = h.services.CreateUser(&user)

	if err != nil {
		h.errorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := createUserResponse{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
	}

	h.writeJSON(w, http.StatusCreated, response)
}

// remove all cookies from the request
// @Summary Logout user
// @Description Remove all cookies from the request
// @Tags auth
// @Produce json
// @Success 200 {object} messageResponse
// @Router /logout [post]
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {

	for _, cookie := range r.Cookies() {
		expiredCookie := &http.Cookie{
			Name:     cookie.Name,
			Value:    "",
			Path:     "/",
			Domain:   cookie.Domain,
			Expires:  time.Unix(0, 0),
			MaxAge:   -1,
			HttpOnly: cookie.HttpOnly,
			Secure:   cookie.Secure,
		}
		http.SetCookie(w, expiredCookie)
	}

	h.writeJSON(w, http.StatusOK, map[string]string{"message": "Logged out"})
}

type loginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Login user and return a cookie with the JWT token and the user info.
// @Summary Login user
// @Description Login user with the provided email and password and return a cookie HttpOnly with the JWT token and the user info.
// @Tags auth
// @Accept json
// @Produce json
// @Param user body loginParams true "User login details"
// @Success 200 {object} loginResponse "Successful login, JWT token set in cookie"
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /login [post]
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var loginParams loginParams

	err := json.NewDecoder(r.Body).Decode(&loginParams)

	if err != nil {
		log.Println("error decoding login params", err)
		h.errorJSON(w, http.StatusBadRequest, "invalid request body")
		return
	}

	fields := map[string]string{
		"email":    loginParams.Email,
		"password": loginParams.Password,
	}

	_, err = helpers.ValidateFields(fields)
	if err != nil {
		h.errorJSON(w, http.StatusBadRequest, "email ou senha inválidos")
		return
	}

	user, err := h.services.GetUserByEmail(loginParams.Email)

	if err != nil {
		log.Println("error getting user by email", err)
		h.errorJSON(w, http.StatusUnauthorized, "email ou senha inválidos")
		return
	}

	isValid := h.comparePasswords(loginParams.Password, user.Password)

	if !isValid {
		log.Println("error invalid password")
		h.errorJSON(w, http.StatusUnauthorized, "email ou senha inválidos")
		return
	}

	token, err := h.JWT.CreateToken(auth.CreateTokenParams{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	})

	if err != nil {
		log.Println("error on login: failed to create token", err)
		h.errorJSON(w, http.StatusInternalServerError, "")
		return
	}

	cookie := h.createHttpOnlyCookie("jwt", token)

	http.SetCookie(w, cookie)

	response := loginResponse{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
	}

	h.writeJSON(w, http.StatusOK, response)
}

type forgotPasswordParams struct {
	Email string `json:"email"`
}

// Send an email with a link to reset the password.
// @Summary Forgot password
// @Description Send an email with a link to reset the password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body forgotPasswordParams true "User email"
// @Success 200 {object} messageResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /forget-password [post]
func (h *Handler) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var params forgotPasswordParams

	err := json.NewDecoder(r.Body).Decode(&params)

	if err != nil {
		log.Println("error decoding forgot password params", err)
		h.errorJSON(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if params.Email == "" {
		log.Println("error missing required fields")
		h.errorJSON(w, http.StatusBadRequest, "Missing required fields")
		return
	}

	user, err := h.services.GetUserByEmail(params.Email)

	if err != nil {
		log.Println("error getting user by email", err)
		h.errorJSON(w, http.StatusNotFound, "Usu´ario não encontrado")
		return
	}

	token, err := h.JWT.CreateToken(auth.CreateTokenParams{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	})

	if err != nil {
		log.Println("error creating token", err)
		h.errorJSON(w, http.StatusInternalServerError, "Failed to create token")
		return
	}

	siteName := os.Getenv("WEBSITE_NAME")
	siteURL := os.Getenv("DOMAIN")
	subject := fmt.Sprintf("%s - Reset password", siteName)

	message := fmt.Sprintf("Click <a href='%s/recover-password?token=%s'>here</a> to reset your password", siteURL, token)

	err = h.services.SendEmail(user.Email, subject, message)

	if err != nil {
		log.Println("error sending reset password email", err)
		h.errorJSON(w, http.StatusInternalServerError, "Failed to send reset password email")
		return
	}

	response := messageResponse{
		Message: "Email sent",
	}

	h.writeJSON(w, http.StatusOK, response)
}

type recoverPasswordParams struct {
	Password string `json:"password"`
	Token    string `json:"token"`
}

// Recover the password with the new password and the token.
// @Summary Recover password
// @Description Recover password with the new password and the token
// @Tags auth
// @Accept json
// @Produce json
// @Param user body recoverPasswordParams true "User new password and token"
// @Success 200 {object} messageResponse
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /recover-password [post]
func (h *Handler) RecoverPassword(w http.ResponseWriter, r *http.Request) {
	var params recoverPasswordParams
	err := json.NewDecoder(r.Body).Decode(&params)

	if err != nil {
		log.Println("error decoding recover password params", err)
		h.errorJSON(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if params.Password == "" || params.Token == "" {
		log.Println("error missing required fields")
		h.errorJSON(w, http.StatusBadRequest, "Missing required fields")
		return
	}

	userID, err := h.JWT.ExtractIDFromToken(params.Token)

	if err != nil {
		log.Println("error extracting user id from token", err)
		h.errorJSON(w, http.StatusUnauthorized, "invalid token")
		return
	}

	hashedPassword, err := h.hashPassword(params.Password)

	if err != nil {
		log.Println("error hashing password", err)
		h.errorJSON(w, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	err = h.services.UpdateUserPasswordByID(&postgresrepo.UpdateUserPasswordByIDParams{
		Password: hashedPassword,
		ID:       userID,
	})

	if err != nil {
		log.Println("error updating password", err)
		h.errorJSON(w, http.StatusInternalServerError, "Failed to update password")
		return
	}

	response := messageResponse{
		Message: "new password set",
	}

	h.writeJSON(w, http.StatusOK, response)
}

type GetUserResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
