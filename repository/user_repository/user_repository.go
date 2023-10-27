package userrepository

import (
	"context"
	"database/sql"

	"github.com/cool-service-go/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur UserRepository) GetUsers(ctx context.Context) ([]model.User, error) {
	users := []model.User{}

	rows, err := ur.db.QueryContext(ctx, GetUsersQuery)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	for rows.Next() {
		var user model.User

		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (ur UserRepository) AddUsers() {}
