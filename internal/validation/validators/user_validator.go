package validators

import (
	"errors"
	"strings"

	"github.com/reangeline/foodscan_backend/internal/dtos"
)

type UserValidator struct{}

func NewUserValidator() *UserValidator {
	return &UserValidator{}
}

func (uv *UserValidator) ValidateUser(user *dtos.CreateUserInput) error {

	if !strings.Contains(user.Email, "@") {
		return errors.New("please add a valid email")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	if user.Name == "" {
		return errors.New("name is required")
	}

	if user.LastName == "" {
		return errors.New("lastname is required")
	}

	return nil
}
