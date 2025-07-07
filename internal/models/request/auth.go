package request

// RegisterRequest описывает структуру запроса для регистрации.
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginRequest описывает структуру запроса для логина.
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LogoutRequest описывает структуру запроса для логаута.
type LogoutRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
