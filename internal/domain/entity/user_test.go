package entity_test

import (
	"testing"

	"github.com/reangeline/foodscan_backend/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyFields_WhenCreateUser_ThenShouldReceiveAnError(t *testing.T) {

	user := entity.User{
		Name:     "",
		LastName: "",
		Email:    "",
	}
	err := user.IsValid()
	assert.Error(t, err)
}

func TestGivenAValidParams_WhenICallNewUser_ThenIShouldReceiveCreateUserWithAllParams(t *testing.T) {

	user := entity.User{
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

	u, err := entity.NewUser("Renato", "Angeline", "reangeline@hotmail.com")
	assert.Nil(t, err)

	assert.Equal(t, "Renato", u.Name)
	assert.Equal(t, "Angeline", u.LastName)
	assert.Equal(t, "reangeline@hotmail.com", u.Email)

}

func TestGivenAInvalidParams_WhenDontSendLastName_ThenIShouldReceiveErrorLastNameIsRequired(t *testing.T) {

	_, err := entity.NewUser("Renato", "", "reangeline@hotmail.com")
	assert.Error(t, err)

}

func TestGivenAInvalidParams_WhenDontSendEmail_ThenIShouldReceiveErrorEmailIsRequired(t *testing.T) {

	_, err := entity.NewUser("Renato", "Angeline", "")
	assert.Error(t, err)

}
