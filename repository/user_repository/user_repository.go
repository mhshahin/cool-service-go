package userrepository

import (
	"context"
	"database/sql"

	"github.com/cool-service-go/model"
	"github.com/lib/pq"
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
