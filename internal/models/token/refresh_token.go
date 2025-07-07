package token

import "time"

// RefreshToken описывает refresh токен пользователя.
type RefreshToken struct {
	ID        int       // Уникальный идентификатор токена
	UserID    int       // ID пользователя
	Token     string    // Строка токена
	ExpiresAt time.Time // Время истечения токена
	CreatedAt time.Time // Время создания токена
}
