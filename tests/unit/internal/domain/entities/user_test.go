package entities_test

import (
	"testing"

	"github.com/reangeline/foodscan_backend/internal/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyFields_WhenCreateUser_ThenShouldReceiveAnError(t *testing.T) {

	user := entities.User{
		Name:     "",
		LastName: "",
		Email:    "",
	}
	err := user.IsValid()
	assert.Error(t, err)
}

func TestGivenAValidParams_WhenICallNewUser_ThenIShouldReceiveCreateUserWithAllParams(t *testing.T) {

	user := entities.User{
		Name:     "Renato",
		LastName: "Angeline",
		Email:    "reangeline@hotmail.com",
	}

	assert.Equal(t, "Renato", user.Name)
	assert.Equal(t, "Angeline", user.LastName)
	assert.Equal(t, "reangeline@hotmail.com", user.Email)

	assert.Nil(t, user.IsValid())

}

func TestGivenAValidParams_WhenICallNewUserFunc_ThenIShouldReceiveCreateUserWithAllParams(t *testing.T) {

	u, err := entities.NewUser("Renato", "Angeline", "reangeline@hotmail.com")
	assert.Nil(t, err)

	assert.Equal(t, "Renato", u.Name)
	assert.Equal(t, "Angeline", u.LastName)
	assert.Equal(t, "reangeline@hotmail.com", u.Email)

}
