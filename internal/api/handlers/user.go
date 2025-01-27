package handlers

import (
	"encoding/json"
	"github.com/gatimugabriel/hotel-reservation-system/internal/config"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/user/entity"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/user/services"
	"github.com/gatimugabriel/hotel-reservation-system/pkg/utils"
	"golang.org/x/oauth2"
	"net/http"
)

type UserHandler struct {
	userService services.UserService
	googleOAuth *oauth2.Config
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
		googleOAuth: config.GoogleOAuthConfig,
	}
}

func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req entity.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, err := h.userService.Create(r.Context(), &req)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to register user")
		return
	}

	utils.RespondJSON(w, http.StatusCreated, user)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req entity.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	accessToken, refreshToken, err := h.userService.Authenticate(r.Context(), &req)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]string{"accessToken": accessToken, "refreshToken": refreshToken})
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {

}