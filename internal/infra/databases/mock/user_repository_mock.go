package mock

import (
	"context"
	"errors"

	"github.com/reangeline/foodscan_backend/internal/domain/entities"
	"github.com/reangeline/foodscan_backend/internal/dtos"
)

type UserRepositoryMock struct {
	Users map[string]*entities.User
}

func NewUserRepositoryMock() *UserRepositoryMock {
	return &UserRepositoryMock{
		Users: make(map[string]*entities.User),
	}
}

func (r *UserRepositoryMock) CreateUser(ctx context.Context, user *entities.User) error {
	r.Users[user.Email] = user
	return nil
}

func (r *UserRepositoryMock) FindByUserEmail(email string) (*dtos.UserOutputDTO, error) {
	user, ok := r.Users[email]

	if !ok {
		return nil, errors.New("email is not found")
	}

	return &dtos.UserOutputDTO{
		IDUser:   user.IDUser,
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
	}, nil
}
