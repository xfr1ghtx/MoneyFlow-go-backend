package service

import (
	"context"
	"errors"

	"github.com/stepanpotapov/moneyflow-go-backend/internal/models/account"
	"github.com/stepanpotapov/moneyflow-go-backend/internal/repository"
)

// BankAccountService реализует бизнес-логику для банковских аккаунтов.
type BankAccountService struct {
	repo *repository.BankAccountRepository // Репозиторий банковских аккаунтов
}

// NewBankAccountService создает новый экземпляр BankAccountService.
func NewBankAccountService(repo *repository.BankAccountRepository) *BankAccountService {
	return &BankAccountService{repo: repo}
}

// Create создает новый банковский аккаунт для пользователя.
func (s *BankAccountService) Create(ctx context.Context, userID int, name string, balance float64, currency string) (*account.BankAccount, error) {
	if name == "" || currency == "" {
		return nil, errors.New("Название и валюта обязательны")
	}
	return s.repo.Create(ctx, userID, name, balance, currency)
}

// Update обновляет банковский аккаунт по id и user_id.
func (s *BankAccountService) Update(ctx context.Context, id, userID int, name string, balance float64, currency string) (*account.BankAccount, error) {
	if name == "" || currency == "" {
		return nil, errors.New("Название и валюта обязательны")
	}
	return s.repo.Update(ctx, id, userID, name, balance, currency)
}

// Delete удаляет банковский аккаунт по id и user_id.
func (s *BankAccountService) Delete(ctx context.Context, id, userID int) error {
	return s.repo.Delete(ctx, id, userID)
}
