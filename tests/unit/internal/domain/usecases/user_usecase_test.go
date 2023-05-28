package usecases_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/reangeline/foodscan_backend/internal/domain/usecases"
	"github.com/reangeline/foodscan_backend/internal/dtos"
	database_test "github.com/reangeline/foodscan_backend/tests/unit/internal/infra/database"
)

func TestUserUseCase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	repo := database_test.NewMemoryUserRepository()
	usecase := usecases.NewUserUseCase(repo)

	input := &dtos.CreateUserInput{
		Name:     "Renato",
		LastName: "Angeline",
		Email:    "reangeline@hotmail.com",
	}

	t.Run("should create user successfully", func(t *testing.T) {
		err := usecase.CreateUser(ctx, input)
		assert.NoError(t, err)
	})

	t.Run("should find user by email", func(t *testing.T) {
		user, err := usecase.FindByEmail(input.Email)
		assert.NoError(t, err)
		assert.Equal(t, user.Email, input.Email)
	})

	t.Run("should not add a email that alrealdy exist", func(t *testing.T) {
		err := usecase.CreateUser(ctx, input)
		assert.Error(t, err)
	})

}
