package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/reangeline/foodscan_backend/internal/presentation/controller"
)

func InitializeUserRoutes(controller *controller.UserController, r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", controller.CreateUserRest)
		r.Get("/", controller.FindUserByEmailRest)
	})
}
