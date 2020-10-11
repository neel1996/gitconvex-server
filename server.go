package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/neel1996/gitconvex-server/global"
	"github.com/neel1996/gitconvex-server/graph"
	"github.com/neel1996/gitconvex-server/graph/generated"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const defaultPort = "9001"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/gitconvexapi", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	router := mux.NewRouter()

	router.Path("/gitconvexapi/graph").Handler(playground.Handler("GraphQL", "/query"))
	router.Handle("/query", srv)
	router.Handle("/gitconvexapi", srv)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./build/")))

	logger := global.Logger{Message: fmt.Sprintf("Gitconvex started on  http://localhost:%v", port)}
	logger.LogInfo()

	log.Fatal(http.ListenAndServe(":"+port, cors.Default().Handler(router)))
}
