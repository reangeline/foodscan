package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/reangeline/foodscan_backend/internal/presentation/controllers"
)

func InitializeUserRoutes(controller *controllers.UserController) chi.Router {
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) {
		r.Post("/", controller.CreateUserRest)
	})

	return r
}
