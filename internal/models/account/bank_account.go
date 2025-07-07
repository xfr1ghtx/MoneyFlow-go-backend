package account

import "time"

// BankAccount описывает банковский аккаунт пользователя.
type BankAccount struct {
	ID        int       // Уникальный идентификатор аккаунта
	UserID    int       // ID пользователя
	Name      string    // Название аккаунта
	Balance   float64   // Баланс
	Currency  string    // Валюта
	CreatedAt time.Time // Дата создания
	UpdatedAt time.Time // Дата обновления
}
