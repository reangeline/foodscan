//go:build wireinject
// +build wireinject

package factory

import (
	"github.com/google/wire"

	"database/sql"

	"github.com/reangeline/foodscan_backend/internal/domain/contract/repository"
	uc_interface "github.com/reangeline/foodscan_backend/internal/domain/contract/usecase"

	"github.com/reangeline/foodscan_backend/internal/domain/usecase"
	"github.com/reangeline/foodscan_backend/internal/infra/database"

	"github.com/reangeline/foodscan_backend/internal/presentation/controller"
	"github.com/reangeline/foodscan_backend/internal/presentation/validation/protocol"
	"github.com/reangeline/foodscan_backend/internal/presentation/validation/validator"
)

var setUserRepositoryDependency = wire.NewSet(
	database.NewUserRepository,
	wire.Bind(new(repository.UserRepositoryInterface), new(*database.UserRepository)),
)

var setUserUseCaseDependency = wire.NewSet(
	usecase.NewUserUseCase,
	wire.Bind(new(uc_interface.UserUseCaseInterface), new(*usecase.UserUseCase)),
)

var setUserValidatorDependency = wire.NewSet(
	validator.NewUserValidator,
	wire.Bind(new(protocol.UserValidatorInterface), new(*validator.UserValidator)),
)

func InitializeUser(db *sql.DB) (*controller.UserController, error) {
	wire.Build(
		setUserRepositoryDependency,
		setUserUseCaseDependency,
		setUserValidatorDependency,
		controller.NewUserController,
	)
	return &controller.UserController{}, nil
}
