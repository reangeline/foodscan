package validators

import (
	"strings"

	"github.com/reangeline/foodscan_backend/internal/dtos"
	err "github.com/reangeline/foodscan_backend/internal/presentation/errors"
)

type UserValidator struct{}

func NewUserValidator() *UserValidator {
	return &UserValidator{}
}

func (uv *UserValidator) ValidateUser(user *dtos.CreateUserInput) error {

	if !strings.Contains(user.Email, "@") {
		return err.ErrValidEmail
	}

	if user.Email == "" {
		return err.ErrEmailIsRequired
	}

	if user.Name == "" {
		return err.ErrNameIsRequired
	}

	if user.LastName == "" {
		return err.ErrLastNameIsRequired
	}

	return nil
}
