package graphql

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/reangeline/foodscan_backend/configs"
	"github.com/reangeline/foodscan_backend/internal/factory"
	"github.com/reangeline/foodscan_backend/internal/infra/graphql/graph"
)

const defaultPort = "8080"

func ServerGQL(db *sql.DB, config *configs.Conf) {

	userController, err := factory.InitializeUser(db)
	if err != nil {
		log.Fatalf("failed to initialize user controller: %v", err)
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		UserController: userController,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.GraphQLServerPort)
	log.Fatal(http.ListenAndServe(":"+config.GraphQLServerPort, nil))
}
