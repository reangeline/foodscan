package entities

import (
	"errors"

	"github.com/reangeline/foodscan_backend/pkg/entities"
)

type User struct {
	IDUser   entities.ID
	Name     string
	LastName string
	Email    string
}

func NewUser(name string, last_name string, email string) (*User, error) {
	user := &User{
		Name:     name,
		LastName: last_name,
		Email:    email,
	}

	user.AddId()

	err := user.IsValid()

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) AddId() {
	u.IDUser = entities.NewID()
}

func (u *User) IsValid() error {

	if u.Name == "" {
		return errors.New("name is required")
	}

	if u.LastName == "" {
		return errors.New("last name is required")
	}

	if u.Email == "" {
		return errors.New("email is required")
	}

	return nil
}
