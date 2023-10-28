package repository

import (
	"database/sql"

	userrepository "github.com/mhshahin/cool-service-go/repository/user_repository"
)

type Repository struct {
	UserRepository userrepository.User
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserRepository: userrepository.NewUserRepository(db),
	}
}
