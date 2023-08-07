package usecases_test

import (
	"context"
	"testing"

	"github.com/reangeline/foodscan_backend/internal/domain/usecases"
	"github.com/reangeline/foodscan_backend/internal/dtos"
	"github.com/reangeline/foodscan_backend/internal/infra/databases/mock"
	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyFields_WhenCreateUser_ThenShouldReceiveAnError(t *testing.T) {

	mock := mock.NewUserRepositoryMock()
	userUseCase := usecases.NewUserUseCase(mock)

	input := &dtos.CreateUserInput{
		Name:     "",
		LastName: "",
		Email:    "",
	}

	err := userUseCase.CreateUser(context.Background(), input)

	assert.Error(t, err)

}

func TestGivenTryAddSameEmail_WhenCreateUser_ThenShouldReceiveAnError(t *testing.T) {

	mock := mock.NewUserRepositoryMock()
	userUseCase := usecases.NewUserUseCase(mock)

	input := &dtos.CreateUserInput{
		Name:     "John",
		LastName: "Doe",
		Email:    "john@example.com",
	}

	err := userUseCase.CreateUser(context.Background(), input)

	assert.NoError(t, err)

	err = userUseCase.CreateUser(context.Background(), input)

	assert.ErrorIs(t, err, usecases.ErrEmailAlreadyExists)

}

func TestGivenValidEmail_WhenFindEmail_ThenShouldReceiveAnUser(t *testing.T) {

	mock := mock.NewUserRepositoryMock()
	userUseCase := usecases.NewUserUseCase(mock)

	input := &dtos.CreateUserInput{
		Name:     "John",
		LastName: "Doe",
		Email:    "john@example.com",
	}

	err := userUseCase.CreateUser(context.Background(), input)

	assert.NoError(t, err)

	user, err := userUseCase.FindByEmail("john@example.com")

	if user.Email != "john@example.com" {
		assert.Error(t, err)
	}

	if user.Name != "John" {
		assert.Error(t, err)
	}

	if user.LastName != "Doe" {
		assert.Error(t, err)
	}

}

func TestGivenEmptyEmail_WhenFindEmail_ThenShouldReceiveAnError(t *testing.T) {
	mock := mock.NewUserRepositoryMock()
	userUseCase := usecases.NewUserUseCase(mock)

	_, err := userUseCase.FindByEmail("john@example.com")

	assert.Error(t, err)

}
