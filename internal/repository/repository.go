package repository

type Repository interface {
	StoreRefreshToken(userID, token string) error
	GetRefreshToken(userID string) (string, error)
	UpdateRefreshToken(userID, token string) error
	ValidateRefreshToken(token string) (string, error)
}
