package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stepanpotapov/moneyflow-go-backend/internal/models/common"
	"github.com/stepanpotapov/moneyflow-go-backend/internal/service"
)

// AuthHandler содержит обработчики HTTP-запросов для аутентификации.
type AuthHandler struct {
	service *service.AuthService // Сервис авторизации
}

// NewAuthHandler создает новый экземпляр AuthHandler.
func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// registerRequest описывает структуру запроса для регистрации.
type registerRequest struct {
	Email    string `json:"email" binding:"required,email"` // Email пользователя
	Password string `json:"password" binding:"required"`    // Пароль пользователя
}

// loginRequest описывает структуру запроса для логина.
type loginRequest struct {
	Email    string `json:"email" binding:"required,email"` // Email пользователя
	Password string `json:"password" binding:"required"`    // Пароль пользователя
}

// logoutRequest описывает структуру запроса для логаута.
type logoutRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"` // Refresh токен
}

// Register обрабатывает запрос на регистрацию пользователя.
// @Summary Регистрация
// @Tags auth
// @Accept json
// @Produce json
// @Param input body registerRequest true "Данные для регистрации"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "ошибка"
// @Router /register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{StatusCode: http.StatusBadRequest, Message: "Некорректные данные"})
		return
	}
	err := h.service.Register(context.Background(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{StatusCode: http.StatusBadRequest, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// Login обрабатывает запрос на вход пользователя.
// @Summary Логин
// @Tags auth
// @Accept json
// @Produce json
// @Param input body loginRequest true "Данные для входа"
// @Success 200 {object} service.Tokens
// @Failure 400 {string} string "ошибка"
// @Router /login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{StatusCode: http.StatusBadRequest, Message: "Некорректные данные"})
		return
	}
	tokens, err := h.service.Login(context.Background(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{StatusCode: http.StatusBadRequest, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, tokens)
}

// Logout обрабатывает запрос на логаут пользователя (инвалидация refresh токена).
// @Summary Логаут
// @Tags auth
// @Accept json
// @Produce json
// @Param input body logoutRequest true "Refresh токен для логаута"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "ошибка"
// @Router /logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	var req logoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{StatusCode: http.StatusBadRequest, Message: "Некорректные данные"})
		return
	}
	err := h.service.Logout(context.Background(), req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{StatusCode: http.StatusBadRequest, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
