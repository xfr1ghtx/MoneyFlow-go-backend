package common

// ErrorResponse описывает структуру ответа с ошибкой.
type ErrorResponse struct {
	StatusCode int    `json:"statusCode"` // HTTP статус ошибки
	Message    string `json:"message"`    // Сообщение с деталями ошибки
}
