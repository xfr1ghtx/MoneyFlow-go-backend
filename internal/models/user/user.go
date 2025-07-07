package user

import "time"

// User представляет пользователя системы.
type User struct {
	ID           int       // Уникальный идентификатор пользователя
	Email        string    // Email пользователя
	PasswordHash string    // Хеш пароля пользователя
	CreatedAt    time.Time // Дата и время создания пользователя
}
