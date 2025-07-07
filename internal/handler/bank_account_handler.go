package handler

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stepanpotapov/moneyflow-go-backend/internal/models/common"
	"github.com/stepanpotapov/moneyflow-go-backend/internal/service"
)

// BankAccountHandler содержит обработчики HTTP-запросов для банковских аккаунтов.
type BankAccountHandler struct {
	service   *service.BankAccountService // Сервис банковских аккаунтов
	jwtSecret string                      // Секрет для JWT
}

// NewBankAccountHandler создает новый экземпляр BankAccountHandler.
func NewBankAccountHandler(service *service.BankAccountService, jwtSecret string) *BankAccountHandler {
	return &BankAccountHandler{service: service, jwtSecret: jwtSecret}
}

// bankAccountRequest описывает структуру запроса для создания/обновления аккаунта.
type bankAccountRequest struct {
	Name     string  `json:"name" binding:"required"`     // Название
	Balance  float64 `json:"balance" binding:"required"`  // Баланс
	Currency string  `json:"currency" binding:"required"` // Валюта
}

// getUserIDFromToken извлекает userID из access token (JWT).
func (h *BankAccountHandler) getUserIDFromToken(c *gin.Context) (int, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return 0, http.ErrNoCookie
	}
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return 0, http.ErrNoCookie
	}
	tokenStr := parts[1]
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.jwtSecret), nil
	})
	if err != nil || !token.Valid {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, http.ErrNoCookie
	}
	userIDf, ok := claims["user_id"].(float64)
	if !ok {
		return 0, http.ErrNoCookie
	}
	return int(userIDf), nil
}

// CreateBankAccount создает новый банковский аккаунт для пользователя.
// @Summary Создать банковский аккаунт
// @Tags accounts
// @Accept json
// @Produce json
// @Param input body request.BankAccountRequest true "Данные аккаунта"
// @Success 200 {object} account.BankAccount
// @Failure 400 {object} common.ErrorResponse "ошибка"
// @Failure 401 {string} string "Неавторизован"
// @Security BearerAuth
// @Router /accounts [post]
func (h *BankAccountHandler) CreateBankAccount(c *gin.Context) {
	userID, err := h.getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, common.ErrorResponse{StatusCode: http.StatusUnauthorized, Message: "Неавторизован"})
		return
	}
	var req bankAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}
	acc, err := h.service.Create(context.Background(), userID, req.Name, req.Balance, req.Currency)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, acc)
}

// UpdateBankAccount обновляет банковский аккаунт по id.
// @Summary Обновить банковский аккаунт
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path int true "ID аккаунта"
// @Param input body bankAccountRequest true "Данные аккаунта"
// @Success 200 {object} account.BankAccount
// @Failure 400 {object} common.ErrorResponse "ошибка"
// @Failure 401 {string} string "Неавторизован"
// @Security BearerAuth
// @Router /accounts/{id} [put]
func (h *BankAccountHandler) UpdateBankAccount(c *gin.Context) {
	userID, err := h.getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, common.ErrorResponse{StatusCode: http.StatusUnauthorized, Message: "Неавторизован"})
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный id"})
		return
	}
	var req bankAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}
	acc, err := h.service.Update(context.Background(), id, userID, req.Name, req.Balance, req.Currency)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, acc)
}

// DeleteBankAccount удаляет банковский аккаунт по id.
// @Summary Удалить банковский аккаунт
// @Tags accounts
// @Param id path int true "ID аккаунта"
// @Success 200 {object} response.MessageResponse
// @Failure 400 {object} common.ErrorResponse "ошибка"
// @Failure 401 {string} string "Неавторизован"
// @Security BearerAuth
// @Router /accounts/{id} [delete]
func (h *BankAccountHandler) DeleteBankAccount(c *gin.Context) {
	userID, err := h.getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, common.ErrorResponse{StatusCode: http.StatusUnauthorized, Message: "Неавторизован"})
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный id"})
		return
	}
	err = h.service.Delete(context.Background(), id, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
