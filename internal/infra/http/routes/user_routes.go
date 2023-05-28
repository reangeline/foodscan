package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/reangeline/foodscan_backend/internal/infra/http/controllers"
)

func InitializeUserRoutes(uc *controllers.UserController) chi.Router {
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) {
		r.Post("/", uc.CreateUser)
	})

	return r
}
