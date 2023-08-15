package validators_test

import (
	"testing"

	"github.com/reangeline/foodscan_backend/internal/dtos"
	"github.com/reangeline/foodscan_backend/internal/presentation/validation/validators"
	"github.com/stretchr/testify/assert"
)

func TestGivenAnInvalidEmail_WhenValidateUser_ThenShouldReciveAnError(t *testing.T) {
	valid := validators.NewUserValidator()

	user := &dtos.CreateUserInput{
		Name:     "abc",
		LastName: "abc",
		Email:    "abc",
	}

	err := valid.ValidateUser(user)

	assert.Error(t, err)

}

func TestGivenAnEmptyName_WhenValidateUser_ThenShouldReciveAnError(t *testing.T) {
	valid := validators.NewUserValidator()

	user := &dtos.CreateUserInput{
		Name:     "",
		LastName: "aaaa",
		Email:    "john@doe.com",
	}

	err := valid.ValidateUser(user)

	assert.Error(t, err)

}

func TestGivenAnEmptyLastName_WhenValidateUser_ThenShouldReciveAnError(t *testing.T) {
	valid := validators.NewUserValidator()

	user := &dtos.CreateUserInput{
		Name:     "aaaa",
		LastName: "",
		Email:    "john@doe.com",
	}

	err := valid.ValidateUser(user)

	assert.Error(t, err)

}

func TestGivenValidFields_WhenValidateUser_ThenShouldNotReciveAnError(t *testing.T) {
	valid := validators.NewUserValidator()

	user := &dtos.CreateUserInput{
		Name:     "aaaa",
		LastName: "xxxx",
		Email:    "john@doe.com",
	}

	err := valid.ValidateUser(user)

	assert.NoError(t, err)

}
