package repository

import (
	"database/sql"

	userrepository "github.com/cool-service-go/repository/user_repository"
)

type Repository struct {
	UserRepository *userrepository.UserRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserRepository: userrepository.NewUserRepository(db),
	}
}
