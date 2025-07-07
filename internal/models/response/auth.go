package response

// TokensResponse содержит access и refresh токены для пользователя.
type TokensResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
