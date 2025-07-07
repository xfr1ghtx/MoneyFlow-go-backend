package service

import (
	"context"
	"errors"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stepanpotapov/moneyflow-go-backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

// AuthService реализует бизнес-логику аутентификации и регистрации пользователей.
type AuthService struct {
	repo        *repository.UserRepository         // Репозиторий пользователей
	refreshRepo *repository.RefreshTokenRepository // Репозиторий refresh токенов
	jwtSecret   string                             // Секрет для подписи JWT
}

// Tokens содержит access и refresh токены для пользователя.
type Tokens struct {
	AccessToken  string `json:"access_token"`  // JWT access token
	RefreshToken string `json:"refresh_token"` // JWT refresh token
}

// NewAuthService создает новый экземпляр AuthService.
func NewAuthService(repo *repository.UserRepository, refreshRepo *repository.RefreshTokenRepository, jwtSecret string) *AuthService {
	return &AuthService{repo: repo, refreshRepo: refreshRepo, jwtSecret: jwtSecret}
}

// Register регистрирует нового пользователя с проверкой сложности пароля и хешированием.
func (s *AuthService) Register(ctx context.Context, email, password string) error {
	if !isPasswordStrong(password) {
		return errors.New("Пароль слишком простой")
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Ошибка при хешировании пароля")
	}
	return s.repo.Create(ctx, email, string(passwordHash))
}

// Login выполняет аутентификацию пользователя по email и паролю, возвращает токены и сохраняет refresh токен в БД.
func (s *AuthService) Login(ctx context.Context, email, password string) (*Tokens, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("Неверный email или пароль")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("Неверный email или пароль")
	}
	access, err := s.generateToken(user.ID, user.Email, 15*time.Minute)
	if err != nil {
		return nil, err
	}
	refresh, err := s.generateToken(user.ID, user.Email, 7*24*time.Hour)
	if err != nil {
		return nil, err
	}
	// Сохраняем refresh токен в БД
	expiresAt := time.Now().Add(7 * 24 * time.Hour)
	err = s.refreshRepo.Save(ctx, user.ID, refresh, expiresAt)
	if err != nil {
		return nil, errors.New("Ошибка сохранения refresh токена")
	}
	return &Tokens{AccessToken: access, RefreshToken: refresh}, nil
}

// Logout удаляет refresh токен из БД (инвалидация токена).
func (s *AuthService) Logout(ctx context.Context, refreshToken string) error {
	return s.refreshRepo.Delete(ctx, refreshToken)
}

// generateToken создает JWT токен с заданным временем жизни.
func (s *AuthService) generateToken(userID int, email string, ttl time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(ttl).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}

// isPasswordStrong проверяет сложность пароля: минимум 8 символов, буквы и цифры.
func isPasswordStrong(password string) bool {
	if len(password) < 8 {
		return false
	}
	matched, _ := regexp.MatchString(`[A-Za-z]`, password)
	if !matched {
		return false
	}
	matched, _ = regexp.MatchString(`[0-9]`, password)
	return matched
}
