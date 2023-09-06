package routes

import (
	"fmt"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/reangeline/foodscan_backend/configs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func InitializeMiddlewares(config *configs.Conf, r chi.Router) {

	r.Use(middleware.Heartbeat("/health"))

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	swaggerUrl := fmt.Sprintf("http://localhost:%s/docs/doc.json", config.WebServerPort)
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(swaggerUrl)))

}
