package http

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/reangeline/foodscan_backend/config"
	"github.com/reangeline/foodscan_backend/internal/factory"
	"github.com/reangeline/foodscan_backend/internal/infra/http/route"
)

func ServerHttp(db *sql.DB, config *config.Conf) {

	iu, err := factory.InitializeUser(db)
	if err != nil {
		log.Fatalf("failed to initialize user controller: %v", err)
	}

	router := chi.NewRouter()

	route.InitializeMiddlewares(config, router)
	route.InitializeUserRoutes(iu, router)

	log.Printf("connect to http://localhost:%s/ for Rest Api", config.WebServerPort)
	err = http.ListenAndServe(":"+config.WebServerPort, router)
	if err != nil {
		panic(err)
	}

}
