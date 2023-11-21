package mock

import (
	"context"
	"errors"

	"github.com/reangeline/foodscan_backend/internal/domain/entity"
	"github.com/reangeline/foodscan_backend/internal/dto"
)

type UserRepositoryMock struct {
	Users map[string]*entity.User
}

func NewUserRepositoryMock() *UserRepositoryMock {
	return &UserRepositoryMock{
		Users: make(map[string]*entity.User),
	}
}

func (r *UserRepositoryMock) CreateUser(ctx context.Context, user *entity.User) error {
	r.Users[user.Email] = user
	return nil
}

func (r *UserRepositoryMock) FindByUserEmail(email string) (*dto.UserOutput, error) {
	user, ok := r.Users[email]

	if !ok {
		return nil, errors.New("email is not found")
	}

	return &dto.UserOutput{
		IDUser:   user.IDUser.String(),
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
	}, nil
}
