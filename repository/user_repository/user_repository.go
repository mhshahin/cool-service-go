package userrepository

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
	"github.com/mhshahin/cool-service-go/model"
)

type User interface {
	GetUsers(ctx context.Context) ([]model.User, error)
	AddUsers(ctx context.Context, users []*model.User) error
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) User {
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
			&user.Email,
			&user.Name,
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

func (ur UserRepository) AddUsers(ctx context.Context, users []*model.User) error {
	tnx, err := ur.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer tnx.Rollback()

	stmt, err := tnx.Prepare(pq.CopyIn("users", "email", "name"))
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, user := range users {
		_, err := stmt.ExecContext(ctx, user.Email, user.Name)
		if err != nil {
			return err
		}
	}

	// Complete the COPY operation within the transaction
	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	err = tnx.Commit()
	if err != nil {
		return err
	}

	return nil
}
