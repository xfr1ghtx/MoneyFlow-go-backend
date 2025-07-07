package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stepanpotapov/moneyflow-go-backend/internal/models/user"
)

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
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	row := r.db.QueryRow(ctx, `SELECT id, email, password_hash, created_at FROM users WHERE email = $1`, email)
	var u user.User
	err := row.Scan(&u.ID, &u.Email, &u.PasswordHash, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
