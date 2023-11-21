package route

import (
	"fmt"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/reangeline/foodscan_backend/config"

	httpSwagger "github.com/swaggo/http-swagger"
)

func InitializeMiddlewares(config *config.Conf, r chi.Router) {

	r.Use(middleware.Heartbeat("/health"))

	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	swaggerUrl := fmt.Sprintf("http://localhost:%s/docs/doc.json", config.WebServerPort)
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(swaggerUrl)))

}
