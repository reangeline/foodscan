package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/reangeline/foodscan_backend/configs"
	"github.com/reangeline/foodscan_backend/internal/infra/http/routes"

	_ "github.com/lib/pq"
	_ "github.com/reangeline/foodscan_backend/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Food Scan API
// @version         1.0
// @description     API
// @termsOfService  http://swagger.io/terms/

// @contact.name   Renato Saraiva Angeline
// @contact.email  reangeline@hotmail.com

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(configs)
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		configs.DBHost, configs.DBPort, configs.DBUser, configs.DBPassword, configs.DBName,
	)

	db, err := sql.Open("postgres", dsn)
	fmt.Println("Postgres connected")

	if err != nil {
		panic(err)
	}

	uc, err := InitializeUserController(db)
	if err != nil {
		log.Fatalf("failed to initialize user controller: %v", err)
	}

	r := routes.InitializeUserRoutes(uc)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	err = http.ListenAndServe(":"+configs.WebServerPort, r)

	if err != nil {
		panic(err)
	}

}
