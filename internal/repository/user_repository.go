package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// User представляет пользователя системы.
type User struct {
	ID           int       // Уникальный идентификатор пользователя
	Email        string    // Email пользователя
	PasswordHash string    // Хеш пароля пользователя
	CreatedAt    time.Time // Дата и время создания пользователя
}

// UserRepository предоставляет методы для работы с пользователями в базе данных.
type UserRepository struct {
	db *pgxpool.Pool // Пул соединений с базой данных
}

// NewUserRepository создает новый экземпляр UserRepository.
func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

// Create добавляет нового пользователя в базу данных.
func (r *UserRepository) Create(ctx context.Context, email, passwordHash string) error {
	_, err := r.db.Exec(ctx, `INSERT INTO users (email, password_hash) VALUES ($1, $2)`, email, passwordHash)
	return err
}

// FindByEmail ищет пользователя по email. Возвращает пользователя или ошибку, если не найден.
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	row := r.db.QueryRow(ctx, `SELECT id, email, password_hash, created_at FROM users WHERE email = $1`, email)
	var u User
	err := row.Scan(&u.ID, &u.Email, &u.PasswordHash, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
