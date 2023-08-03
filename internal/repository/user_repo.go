package repository

import (
	"crud_mysql_api_auth/internal/models"

	"github.com/rs/zerolog/log"
)

type UserRepository interface {
	ReadUser() ([]models.User, error)
	GetUserByUsername(username string) (*models.User, error)
}

func (r *RepositoryImpl) ReadUser() ([]models.User, error) {
	query := "SELECT * FROM users"

	var users []models.User
	err := r.DB.Read.Select(&users, query)
	if err != nil {
		log.Error().Err(err).Msg("Something went wrong")
		return nil, err
	}
	return users, nil
}

func (r *RepositoryImpl) GetUserByUsername(username string) (*models.User, error) {
	query := "SELECT * FROM users WHERE username = ?"

	var user models.User
	err := r.DB.Read.Get(&user, query, username)
	if err != nil {
		log.Error().Err(err).Msg("Something went wrong")
		return nil, err
	}
	return &user, nil
}
