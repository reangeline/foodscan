package repository

import (
	"context"

	"github.com/reangeline/foodscan_backend/internal/domain/entity"
	"github.com/reangeline/foodscan_backend/internal/dto"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *entity.User) error
	FindByUserEmail(email string) (*dto.UserOutput, error)
}
