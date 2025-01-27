package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gatimugabriel/hotel-reservation-system/internal/config"
	"github.com/gatimugabriel/hotel-reservation-system/internal/constants"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/user/entity"
	"github.com/gatimugabriel/hotel-reservation-system/internal/domain/user/services"
	"github.com/gatimugabriel/hotel-reservation-system/pkg/utils"
	"github.com/gatimugabriel/hotel-reservation-system/pkg/utils/input"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/oauth2"
	"net/http"
	"regexp"
	"time"
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
	var req entity.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	//validate & sanitize inputs
	req.FirstName = input.SanitizeString(req.FirstName)
	req.LastName = input.SanitizeString(req.LastName)
	req.Email = input.SanitizeString(req.Email)
	req.Phone = input.SanitizeString(req.Phone)
	if req.Role == "" {
		req.Role = constants.GUEST
	}

	if validationErrors := input.Validate(&req); validationErrors != nil {
		utils.RespondJSON(w, http.StatusBadRequest, validationErrors)
		return
	}

	user, err := h.userService.Create(r.Context(), &req)
	if err != nil {
		errorMessage := ""
		errorStatus := http.StatusInternalServerError

		// Check for unique constraint violation
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				//  Extract column name and value for the detail message
				re := regexp.MustCompile(`Key \(([^)]+)\)=\([^)]*\) already exists.`)
				matches := re.FindStringSubmatch(pgErr.Detail)
				if len(matches) == 2 {
					columnName := matches[1]
					errorMessage = fmt.Sprintf("%s already exists", columnName)
					errorStatus = http.StatusConflict
				} else {
					errorMessage = fmt.Sprintf("%s already exists", pgErr.ConstraintName)
					errorStatus = http.StatusConflict
				}

				utils.RespondError(w, errorStatus, errorMessage)
				return
			}
		}

		// default
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// generate tokens for new user
	accessToken, refreshToken, _ := utils.GenerateTokens(user.ID.String(), user.Role)

	// Set HTTP_only cookies
	http.SetCookie(w, &http.Cookie{
		Name:     "accessToken",
		Value:    accessToken,
		Expires:  time.Now().Add(15 * time.Minute),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		Expires:  time.Now().Add(30 * 24 * time.Hour),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	utils.RespondJSON(w, http.StatusCreated, map[string]interface{}{"user": user, "accessToken": accessToken, "refreshToken": refreshToken})
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req entity.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	req.Email = input.SanitizeString(req.Email)

	accessToken, refreshToken, err := h.userService.Authenticate(r.Context(), &req)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Set HTTP_only cookies
	http.SetCookie(w, &http.Cookie{
		Name:     "accessToken",
		Value:    accessToken,
		Expires:  time.Now().Add(15 * time.Minute),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		Expires:  time.Now().Add(30 * 24 * time.Hour),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	utils.RespondJSON(w, http.StatusOK, map[string]string{"accessToken": accessToken, "refreshToken": refreshToken})
}

func (h *UserHandler) GetUserProfile(w http.ResponseWriter, r *http.Request) {

}