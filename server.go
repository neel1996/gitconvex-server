package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/neel1996/gitconvex-server/global"
	"github.com/neel1996/gitconvex-server/graph"
	"github.com/neel1996/gitconvex-server/graph/generated"
	"github.com/neel1996/gitconvex-server/utils"
	"github.com/rs/cors"
	"log"
	"net/http"
)

const defaultPort = "9001"

var (
	Port string
)

func main() {
	//var envConfig *utils.EnvConfig
	if err := utils.EnvConfigValidator(); err != nil {
		_ = utils.EnvConfigFileGenerator()
	} else {
		if err := utils.EnvConfigFileGenerator(); err == nil {
			envConfig := *utils.EnvConfigFileReader()
			Port = envConfig.Port
		}
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/gitconvexapi", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	router := mux.NewRouter()

	router.Path("/gitconvexapi/graph").Handler(playground.Handler("GraphQL", "/query"))
	router.Handle("/query", srv)
	router.Handle("/gitconvexapi", srv)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./build/")))

	logger := global.Logger{Message: fmt.Sprintf("Gitconvex started on  http://localhost:%v", Port)}
	logger.LogInfo()

	if Port != "" {
		log.Fatal(http.ListenAndServe(":"+Port, cors.Default().Handler(router)))
	} else {
		log.Fatal(http.ListenAndServe(":"+defaultPort, cors.Default().Handler(router)))
	}
}
