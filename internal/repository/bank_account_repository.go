package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stepanpotapov/moneyflow-go-backend/internal/models/account"
)

// BankAccountRepository предоставляет методы для работы с банковскими аккаунтами в БД.
type BankAccountRepository struct {
	db *pgxpool.Pool // Пул соединений с БД
}

// NewBankAccountRepository создает новый экземпляр BankAccountRepository.
func NewBankAccountRepository(db *pgxpool.Pool) *BankAccountRepository {
	return &BankAccountRepository{db: db}
}

// Create создает новый банковский аккаунт для пользователя.
func (r *BankAccountRepository) Create(ctx context.Context, userID int, name string, balance float64, currency string) (*account.BankAccount, error) {
	row := r.db.QueryRow(ctx, `INSERT INTO bank_accounts (user_id, name, balance, currency) VALUES ($1, $2, $3, $4) RETURNING id, user_id, name, balance, currency, created_at, updated_at`, userID, name, balance, currency)
	var acc account.BankAccount
	err := row.Scan(&acc.ID, &acc.UserID, &acc.Name, &acc.Balance, &acc.Currency, &acc.CreatedAt, &acc.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &acc, nil
}

// Update обновляет банковский аккаунт по id и user_id.
func (r *BankAccountRepository) Update(ctx context.Context, id, userID int, name string, balance float64, currency string) (*account.BankAccount, error) {
	row := r.db.QueryRow(ctx, `UPDATE bank_accounts SET name=$1, balance=$2, currency=$3, updated_at=NOW() WHERE id=$4 AND user_id=$5 RETURNING id, user_id, name, balance, currency, created_at, updated_at`, name, balance, currency, id, userID)
	var acc account.BankAccount
	err := row.Scan(&acc.ID, &acc.UserID, &acc.Name, &acc.Balance, &acc.Currency, &acc.CreatedAt, &acc.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &acc, nil
}

// Delete удаляет банковский аккаунт по id и user_id.
func (r *BankAccountRepository) Delete(ctx context.Context, id, userID int) error {
	_, err := r.db.Exec(ctx, `DELETE FROM bank_accounts WHERE id=$1 AND user_id=$2`, id, userID)
	return err
}
