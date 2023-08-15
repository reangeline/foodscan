package controllers_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/reangeline/foodscan_backend/internal/domain/contracts/usecases"
	"github.com/reangeline/foodscan_backend/internal/dtos"
	"github.com/reangeline/foodscan_backend/internal/presentation/controllers"
	"github.com/stretchr/testify/assert"
)

// MockUserUseCase is a mock implementation of usecases.UserUseCaseInterface
type MockUserUseCase struct {
	usecases.UserUseCaseInterface
}

func (m *MockUserUseCase) CreateUser(ctx context.Context, user *dtos.CreateUserInput) error {
	return nil
}

// MockUserValidator is a mock implementation of protocols.UserValidatorInterface
type MockUserValidator struct{}

func (m *MockUserValidator) ValidateUser(user *dtos.CreateUserInput) error {
	return nil
}

func TestUserController_CreateUser(t *testing.T) {

	userUseCase := &MockUserUseCase{}
	userValidator := &MockUserValidator{}

	userController := controllers.NewUserController(userUseCase, userValidator)

	t.Run("Successful user creation", func(t *testing.T) {
		requestBody := `{"name": "John", "email": "john@example.com"}`

		req, err := http.NewRequest("POST", "/users", bytes.NewReader([]byte(requestBody)))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		userController.CreateUser(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
	})

	t.Run("Invalid request body", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/users", bytes.NewReader([]byte("invalid-json")))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()

		userController.CreateUser(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)

		// var responseBody Error
		// err = json.Unmarshal(rr.Body.Bytes(), &responseBody)
		// assert.NoError(t, err)
		// assert.Equal(t, "invalid character 'i' looking for beginning of value", responseBody.Message)
	})

}
