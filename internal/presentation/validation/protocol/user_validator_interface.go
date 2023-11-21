package protocol

import "github.com/reangeline/foodscan_backend/internal/dto"

type UserValidatorInterface interface {
	ValidateUser(user *dto.UserInput) error
	ValidateUserEmail(email string) error
}
