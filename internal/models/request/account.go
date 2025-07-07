package request

// BankAccountRequest описывает структуру запроса для создания/обновления банковского аккаунта.
type BankAccountRequest struct {
	Name     string  `json:"name" binding:"required"`
	Balance  float64 `json:"balance" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
}
