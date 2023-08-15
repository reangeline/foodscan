package usecases

import (
	"context"

	"github.com/reangeline/foodscan_backend/internal/dtos"
)

type UserUseCaseInterface interface {
	CreateUser(ctx context.Context, input *dtos.CreateUserInput) error
	CheckEmailExists(email string) (bool, error)
	FindByEmail(email string) (*dtos.UserOutputDTO, error)
}
