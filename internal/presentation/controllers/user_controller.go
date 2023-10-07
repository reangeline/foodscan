package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/reangeline/foodscan_backend/internal/domain/contracts/usecases"
	"github.com/reangeline/foodscan_backend/internal/dtos"
	"github.com/reangeline/foodscan_backend/internal/infra/graphql/graph/model"
	"github.com/reangeline/foodscan_backend/internal/presentation/validation/protocols"
)

type Error struct {
	Message string `json:"message"`
}

type UserController struct {
	userUseCase   usecases.UserUseCaseInterface
	userValidator protocols.UserValidatorInterface
}

func NewUserController(
	userUseCase usecases.UserUseCaseInterface,
	userValidator protocols.UserValidatorInterface,
) *UserController {
	return &UserController{
		userUseCase:   userUseCase,
		userValidator: userValidator,
	}
}

// Create user godoc
// @Summary      Create user
// @Description  Create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request     body      dtos.CreateUserInput  true  "user request"
// @Success      201
// @Failure      500         {object}  Error
// @Router       /users [post]
func (u *UserController) CreateUserRest(w http.ResponseWriter, r *http.Request) {
	var user dtos.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	ctx := r.Context()
	err = u.CreateUser(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (u *UserController) CreateUserGraphql(ctx context.Context, input model.NewUser) error {
	user := dtos.CreateUserInput{
		Name:     input.Name,
		LastName: input.LastName,
		Email:    input.Email,
	}

	err := u.CreateUser(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserController) CreateUser(ctx context.Context, input dtos.CreateUserInput) error {

	err := u.userValidator.ValidateUser(&input)

	if err != nil {
		return err
	}

	err = u.userUseCase.CreateUser(ctx, &input)

	if err != nil {
		return err
	}

	return nil
}

// @Summary      Find user by email
// @Description  Find user by email
// @Tags         find users
// @Accept       json
// @Produce      json
// @Param        email query string true "Endereço de email do usuário"
// @Success      200 {object} dtos.UserOutputDTO
// @Failure      400 {object} Error
// @Router       /users [get]
func (u *UserController) FindUserByEmailRest(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	ctx := r.Context()

	user, err := u.FindUserByEmail(ctx, email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func (u *UserController) FindUserByEmailGraphql(ctx context.Context, email string) (*dtos.UserOutputDTO, error) {
	user, err := u.FindUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserController) FindUserByEmail(ctx context.Context, email string) (*dtos.UserOutputDTO, error) {
	err := u.userValidator.ValidateUserEmail(email)

	if err != nil {
		return nil, err
	}

	user, err := u.userUseCase.FindUserByEmail(ctx, email)

	if err != nil {
		return user, nil
	}

	return user, nil
}
