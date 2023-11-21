package usecase_test

import (
	"context"
	"testing"

	"github.com/reangeline/foodscan_backend/internal/domain/usecase"
	"github.com/reangeline/foodscan_backend/internal/dto"
	"github.com/reangeline/foodscan_backend/internal/infra/database/mock"
	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyFields_WhenCreateUser_ThenShouldReceiveAnError(t *testing.T) {

	mock := mock.NewUserRepositoryMock()
	userUseCase := usecase.NewUserUseCase(mock)

	input := &dto.UserInput{
		Name:     "",
		LastName: "",
		Email:    "",
	}

	err := userUseCase.CreateUser(context.Background(), input)

	assert.Error(t, err)

}

func TestGivenTryAddSameEmail_WhenCreateUser_ThenShouldReceiveAnError(t *testing.T) {

	mock := mock.NewUserRepositoryMock()
	userUseCase := usecase.NewUserUseCase(mock)

	input := &dto.UserInput{
		Name:     "John",
		LastName: "Doe",
		Email:    "john@example.com",
	}

	err := userUseCase.CreateUser(context.Background(), input)

	assert.NoError(t, err)

	err = userUseCase.CreateUser(context.Background(), input)

	assert.ErrorIs(t, err, usecase.ErrEmailAlreadyExists)

}

func TestGivenValidEmail_WhenFindEmail_ThenShouldReceiveAnUser(t *testing.T) {

	mock := mock.NewUserRepositoryMock()
	userUseCase := usecase.NewUserUseCase(mock)

	input := &dto.UserInput{
		Name:     "John",
		LastName: "Doe",
		Email:    "john@example.com",
	}

	err := userUseCase.CreateUser(context.Background(), input)

	assert.NoError(t, err)

	user, err := userUseCase.FindUserByEmail(context.Background(), "john@example.com")

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
	userUseCase := usecase.NewUserUseCase(mock)

	_, err := userUseCase.FindUserByEmail(context.Background(), "john@example.com")

	assert.Error(t, err)

}
