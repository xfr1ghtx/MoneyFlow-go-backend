package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stepanpotapov/moneyflow-go-backend/internal/models/token"
)

// RefreshTokenRepository предоставляет методы для работы с refresh токенами в БД.
type RefreshTokenRepository struct {
	db *pgxpool.Pool // Пул соединений с БД
}

// NewRefreshTokenRepository создает новый экземпляр RefreshTokenRepository.
func NewRefreshTokenRepository(db *pgxpool.Pool) *RefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}

// Save сохраняет refresh токен в базе данных.
func (r *RefreshTokenRepository) Save(ctx context.Context, userID int, tokenStr string, expiresAt, createdAt time.Time) error {
	_, err := r.db.Exec(ctx, `INSERT INTO refresh_tokens (user_id, token, expires_at, created_at) VALUES ($1, $2, $3, $4)`, userID, tokenStr, expiresAt, createdAt)
	return err
}

// Delete удаляет refresh токен из базы данных по значению токена.
func (r *RefreshTokenRepository) Delete(ctx context.Context, tokenStr string) error {
	_, err := r.db.Exec(ctx, `DELETE FROM refresh_tokens WHERE token = $1`, tokenStr)
	return err
}

// FindByToken ищет refresh токен по значению токена.
func (r *RefreshTokenRepository) FindByToken(ctx context.Context, tokenStr string) (*token.RefreshToken, error) {
	row := r.db.QueryRow(ctx, `SELECT id, user_id, token, expires_at, created_at FROM refresh_tokens WHERE token = $1`, tokenStr)
	var rt token.RefreshToken
	err := row.Scan(&rt.ID, &rt.UserID, &rt.Token, &rt.ExpiresAt, &rt.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &rt, nil
}
