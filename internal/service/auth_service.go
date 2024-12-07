package service

import (
	"errors"
	"test-medods/internal/models"
	"test-medods/internal/repository"
	"test-medods/internal/utils"
)

type AuthService struct {
	repo repository.Repository
}

func NewAuthService(repo repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) GenerateTokenPair(userID, ip string) (*models.TokenPair, error) {
	accessToken, err := utils.GenerateJWT(userID, ip)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	err = s.repo.StoreRefreshToken(userID, refreshToken)
	if err != nil {
		return nil, err
	}

	return &models.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) RefreshTokenPair(refreshToken, ip string) (*models.TokenPair, error) {
	userID, err := s.repo.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	storedToken, err := s.repo.GetRefreshToken(userID)
	if err != nil {
		return nil, err
	}

	if !utils.CompareHashAndToken(storedToken, refreshToken) {
		return nil, errors.New("invalid refresh token")
	}

	accessToken, err := utils.GenerateJWT(userID, ip)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	err = s.repo.UpdateRefreshToken(userID, newRefreshToken)
	if err != nil {
		return nil, err
	}

	return &models.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
