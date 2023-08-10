package protocols

import "github.com/reangeline/foodscan_backend/internal/dtos"

type UserValidatorInterface interface {
	ValidateUser(user *dtos.CreateUserInput) error
}
