package usecase

import (
	"context"

	"github.com/reangeline/foodscan_backend/internal/dto"
)

type UserUseCaseInterface interface {
	CreateUser(ctx context.Context, input *dto.UserInput) error
	CheckEmailExists(email string) (bool, error)
	FindUserByEmail(ctx context.Context, email string) (*dto.UserOutput, error)
}
