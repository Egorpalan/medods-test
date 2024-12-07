package repository

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(dbURL string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db: db}, nil
}

func (r *PostgresRepository) StoreRefreshToken(userID, token string) error {
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = r.db.Exec("INSERT INTO refresh_tokens (user_id, token) VALUES ($1, $2)", userID, hashedToken)
	return err
}

func (r *PostgresRepository) GetRefreshToken(userID string) (string, error) {
	var hashedToken string
	err := r.db.QueryRow("SELECT token FROM refresh_tokens WHERE user_id = $1", userID).Scan(&hashedToken)
	if err != nil {
		return "", err
	}
	return hashedToken, nil
}

func (r *PostgresRepository) UpdateRefreshToken(userID, token string) error {
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = r.db.Exec("UPDATE refresh_tokens SET token = $1 WHERE user_id = $2", hashedToken, userID)
	return err
}

func (r *PostgresRepository) ValidateRefreshToken(token string) (string, error) {
	var userID string
	var hashedToken string
	err := r.db.QueryRow("SELECT user_id, token FROM refresh_tokens WHERE token = $1", token).Scan(&userID, &hashedToken)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hashedToken), []byte(token)); err != nil {
		return "", errors.New("invalid refresh token")
	}
	return userID, nil
}
