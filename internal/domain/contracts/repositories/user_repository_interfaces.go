package repositories

import (
	"context"

	"github.com/reangeline/foodscan_backend/internal/domain/entities"
	"github.com/reangeline/foodscan_backend/internal/dtos"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *entities.User) error
	FindByUserEmail(email string) (*dtos.UserOutputDTO, error)
}
