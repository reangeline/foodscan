package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/reangeline/foodscan_backend/internal/domain/contracts/usecases"
	"github.com/reangeline/foodscan_backend/internal/dtos"
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
func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dtos.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = u.userValidator.ValidateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	err = u.userUseCase.CreateUser(ctx, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
