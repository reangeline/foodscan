package databases

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/reangeline/foodscan_backend/internal/domain/entities"
	"github.com/reangeline/foodscan_backend/internal/dtos"
	"github.com/reangeline/foodscan_backend/internal/infra/databases/postgres/sqlc"
)

type UserRepository struct {
	dbConn *sql.DB
	sqlc   *sqlc.Queries
}

func NewUserRepository(dbConn *sql.DB) *UserRepository {
	return &UserRepository{
		dbConn: dbConn,
		sqlc:   sqlc.New(dbConn),
	}
}

func (u *UserRepository) callTx(ctx context.Context, fn func(*sqlc.Queries) error) error {
	tx, err := u.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := sqlc.New(tx)
	err = fn(q)
	if err != nil {
		if errRb := tx.Rollback(); errRb != nil {
			return fmt.Errorf("error on rollback: %v, original error: %w", errRb, err)
		}
		return err
	}
	return tx.Commit()
}

func (u *UserRepository) CreateUser(ctx context.Context, user *entities.User) error {

	err := u.callTx(ctx, func(q *sqlc.Queries) error {
		var err error

		// profilePicture := sql.NullString{String: user.ProfilePicture, Valid: false}

		err = q.CreateUser(ctx, sqlc.CreateUserParams{
			IDUser:   user.IDUser,
			Name:     user.Name,
			LastName: user.LastName,
			Email:    user.Email,
		})
		if err != nil {
			return err
		}

		return nil

	})

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (u *UserRepository) FindByEmail(email string) (*dtos.UserOutputDTO, error) {
	ctx := context.Background()
	user, err := u.sqlc.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &dtos.UserOutputDTO{
		IDUser:   user.IDUser,
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
	}, nil
}
