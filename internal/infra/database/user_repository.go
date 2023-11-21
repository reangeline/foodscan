package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/reangeline/foodscan_backend/internal/domain/entity"
	dtos "github.com/reangeline/foodscan_backend/internal/dto"
	"github.com/reangeline/foodscan_backend/internal/infra/database/postgres/sqlc"
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

func (u *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {

	err := u.callTx(ctx, func(q *sqlc.Queries) error {

		err := q.CreateUser(ctx, sqlc.CreateUserParams{
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

func (u *UserRepository) FindByUserEmail(email string) (*dtos.UserOutput, error) {
	ctx := context.Background()
	user, err := u.sqlc.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	var createdAt time.Time
	if user.CreatedAt.Valid {
		createdAt = user.CreatedAt.Time
	}

	return &dtos.UserOutput{
		IDUser:    user.IDUser.String(),
		Name:      user.Name,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: createdAt,
	}, nil
}
