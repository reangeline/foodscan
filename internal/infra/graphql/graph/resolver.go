package graph

import (
	controller "github.com/reangeline/foodscan_backend/internal/presentation/controller"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserController *controller.UserController
}
