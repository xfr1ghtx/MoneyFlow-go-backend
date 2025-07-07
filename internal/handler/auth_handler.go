package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stepanpotapov/moneyflow-go-backend/internal/models/common"
	req "github.com/stepanpotapov/moneyflow-go-backend/internal/models/request"
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

// Register обрабатывает запрос на регистрацию пользователя.
// @Summary Регистрация
// @Tags auth
// @Accept json
// @Produce json
// @Param input body request.RegisterRequest true "Данные для регистрации"
// @Success 200 {object} response.MessageResponse
// @Failure 400 {object} common.ErrorResponse "ошибка"
// @Router /register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var reqBody req.RegisterRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{StatusCode: http.StatusBadRequest, Message: "Некорректные данные"})
		return
	}
	err := h.service.Register(context.Background(), reqBody.Email, reqBody.Password)
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
// @Param input body request.LoginRequest true "Данные для входа"
// @Success 200 {object} response.TokensResponse
// @Failure 400 {object} common.ErrorResponse "ошибка"
// @Router /login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var reqBody req.LoginRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{StatusCode: http.StatusBadRequest, Message: "Некорректные данные"})
		return
	}
	tokens, err := h.service.Login(context.Background(), reqBody.Email, reqBody.Password)
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
// @Param input body request.LogoutRequest true "Refresh токен для логаута"
// @Success 200 {object} response.MessageResponse
// @Failure 400 {object} common.ErrorResponse "ошибка"
// @Router /logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	var reqBody req.LogoutRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{StatusCode: http.StatusBadRequest, Message: "Некорректные данные"})
		return
	}
	err := h.service.Logout(context.Background(), reqBody.RefreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{StatusCode: http.StatusBadRequest, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
