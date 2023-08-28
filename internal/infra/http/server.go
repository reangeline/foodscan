package http

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/reangeline/foodscan_backend/configs"
	"github.com/reangeline/foodscan_backend/internal/factory"
	"github.com/reangeline/foodscan_backend/internal/infra/http/routes"
)

func ServerHttp(db *sql.DB, config *configs.Conf) {

	iu, err := factory.InitializeUser(db)
	if err != nil {
		log.Fatalf("failed to initialize user controller: %v", err)
	}

	r := routes.InitializeUserRoutes(iu)

	swaggerUrl := fmt.Sprintf("http://localhost:%s/docs/doc.json", config.WebServerPort)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(swaggerUrl)))

	r.Get("/teste", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, config.DBUser)
	})

	log.Printf("connect to http://localhost:%s/ for Rest Api", config.WebServerPort)

	err = http.ListenAndServe(":"+config.WebServerPort, r)
	if err != nil {
		panic(err)
	}

}
