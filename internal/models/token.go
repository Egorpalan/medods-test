package models

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenRequest struct {
	UserID string `json:"user_id"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}
