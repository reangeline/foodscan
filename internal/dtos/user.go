package dtos

import (
	"time"

	"github.com/reangeline/foodscan_backend/pkg/entities"
)

type CreateUserInput struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
}

type UserOutputDTO struct {
	IDUser    entities.ID `json:"id"`
	Name      string      `json:"name"`
	LastName  string      `json:"last_name"`
	Email     string      `json:"email"`
	CreatedAt time.Time   `json:"created_at"`
}
