package service

import (
	"test-medods/internal/models"
	"test-medods/internal/utils"
)

type TokenService struct {
}

func NewTokenService() *TokenService {
	return &TokenService{}
}

func (s *TokenService) GenerateTokenPair(userID, ip string) (*models.TokenPair, error) {
	accessToken, err := utils.GenerateJWT(userID, ip)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	return &models.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
