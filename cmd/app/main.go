package main

import (
	"database/sql"
	"fmt"

	"github.com/reangeline/foodscan_backend/config"
	"github.com/reangeline/foodscan_backend/internal/infra/http"

	_ "github.com/lib/pq"
	_ "github.com/reangeline/foodscan_backend/doc"

	"github.com/reangeline/foodscan_backend/internal/infra/graphql"
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
	configs, err := config.LoadConfig(".")
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

	go graphql.ServerGQL(db, configs)

	http.ServerHttp(db, configs)

}
