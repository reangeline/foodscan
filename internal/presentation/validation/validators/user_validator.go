package validators

import (
	"strings"

	"github.com/reangeline/foodscan_backend/internal/dtos"
	errors "github.com/reangeline/foodscan_backend/internal/presentation/errors"
)

type UserValidator struct {
}

func NewUserValidator() *UserValidator {
	return &UserValidator{}
}

func (uv *UserValidator) ValidateUser(user *dtos.CreateUserInput) error {

	err := uv.ValidateUserEmail(user.Email)
	if err != nil {
		return err
	}

	if user.Name == "" {
		return errors.ErrNameIsRequired
	}

	if user.LastName == "" {
		return errors.ErrLastNameIsRequired
	}

	return nil
}

func (uv *UserValidator) ValidateUserEmail(email string) error {

	if !strings.Contains(email, "@") {
		return errors.ErrValidEmail
	}

	return nil
}
