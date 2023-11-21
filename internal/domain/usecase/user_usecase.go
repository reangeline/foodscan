package usecase

import (
	"context"
	"errors"

	"github.com/reangeline/foodscan_backend/internal/domain/contract/repository"
	"github.com/reangeline/foodscan_backend/internal/domain/entity"
	"github.com/reangeline/foodscan_backend/internal/dto"
)

type UserUseCase struct {
	userRepository repository.UserRepositoryInterface
}

func NewUserUseCase(userRepository repository.UserRepositoryInterface) *UserUseCase {
	return &UserUseCase{userRepository}
}

var (
	ErrEmailAlreadyExists = errors.New("email already exist")
)

func (u *UserUseCase) CreateUser(ctx context.Context, input *dto.UserInput) error {

	isExist, _ := u.CheckEmailExists(input.Email)

	if isExist {
		return ErrEmailAlreadyExists
	}

	user, err := entity.NewUser(input.Name, input.LastName, input.Email)
	if err != nil {
		return err
	}

	if err := u.userRepository.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func (u *UserUseCase) CheckEmailExists(email string) (bool, error) {
	_, err := u.userRepository.FindByUserEmail(email)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (u *UserUseCase) FindUserByEmail(ctx context.Context, email string) (*dto.UserOutput, error) {
	user, err := u.userRepository.FindByUserEmail(email)

	if err != nil {
		return nil, err
	}

	return user, nil

}
