//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	"database/sql"

	"github.com/reangeline/foodscan_backend/internal/domain/contracts/repositories"
	uc_interface "github.com/reangeline/foodscan_backend/internal/domain/contracts/usecases"

	"github.com/reangeline/foodscan_backend/internal/domain/usecases"
	"github.com/reangeline/foodscan_backend/internal/infra/databases"

	"github.com/reangeline/foodscan_backend/internal/presentation/controllers"
	"github.com/reangeline/foodscan_backend/internal/presentation/validation/protocols"
	"github.com/reangeline/foodscan_backend/internal/presentation/validation/validators"
)

var setUserRepositoryDependency = wire.NewSet(
	databases.NewUserRepository,
	wire.Bind(new(repositories.UserRepositoryInterface), new(*databases.UserRepository)),
)

var setUserUseCaseDependency = wire.NewSet(
	usecases.NewUserUseCase,
	wire.Bind(new(uc_interface.UserUseCaseInterface), new(*usecases.UserUseCase)),
)

var setUserValidatorDependency = wire.NewSet(
	validators.NewUserValidator,
	wire.Bind(new(protocols.UserValidatorInterface), new(*validators.UserValidator)),
)

func InitializeUserController(db *sql.DB) (*controllers.UserController, error) {
	wire.Build(
		setUserRepositoryDependency,
		setUserUseCaseDependency,
		setUserValidatorDependency,
		controllers.NewUserController,
	)
	return &controllers.UserController{}, nil
}
