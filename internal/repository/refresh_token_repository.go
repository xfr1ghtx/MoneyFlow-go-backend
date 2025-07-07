package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// RefreshToken описывает refresh токен пользователя.
type RefreshToken struct {
	ID        int       // Уникальный идентификатор токена
	UserID    int       // ID пользователя
	Token     string    // Строка токена
	ExpiresAt time.Time // Время истечения токена
	CreatedAt time.Time // Время создания токена
}

// RefreshTokenRepository предоставляет методы для работы с refresh токенами в БД.
type RefreshTokenRepository struct {
	db *pgxpool.Pool // Пул соединений с БД
}

// NewRefreshTokenRepository создает новый экземпляр RefreshTokenRepository.
func NewRefreshTokenRepository(db *pgxpool.Pool) *RefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}

// Save сохраняет refresh токен в базе данных.
func (r *RefreshTokenRepository) Save(ctx context.Context, userID int, token string, expiresAt time.Time) error {
	_, err := r.db.Exec(ctx, `INSERT INTO refresh_tokens (user_id, token, expires_at) VALUES ($1, $2, $3)`, userID, token, expiresAt)
	return err
}

// Delete удаляет refresh токен из базы данных по значению токена.
func (r *RefreshTokenRepository) Delete(ctx context.Context, token string) error {
	_, err := r.db.Exec(ctx, `DELETE FROM refresh_tokens WHERE token = $1`, token)
	return err
}

// FindByToken ищет refresh токен по значению токена.
func (r *RefreshTokenRepository) FindByToken(ctx context.Context, token string) (*RefreshToken, error) {
	row := r.db.QueryRow(ctx, `SELECT id, user_id, token, expires_at, created_at FROM refresh_tokens WHERE token = $1`, token)
	var rt RefreshToken
	err := row.Scan(&rt.ID, &rt.UserID, &rt.Token, &rt.ExpiresAt, &rt.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &rt, nil
}
