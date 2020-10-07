package main

import (
	"encoding/json"
	"fmt"
	"github.com/neel1996/gitconvex-server/api"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/neel1996/gitconvex-server/model"
)

const defaultPort = "9002"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := mux.NewRouter()

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./build/")))

	api.HealthCheckApi()

	router.HandleFunc("/gitconvexapi", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.Header().Set("Accept", "application/json")

		err := json.NewEncoder(w).Encode(model.StatusResponseModel{Status: "Completed", Message: "Request processed successfull!"})

		if err != nil {
			fmt.Printf("Error occurred : %v", err)
			panic(err)
		}

	}).Methods("GET")

	port = ":" + port
	err := http.ListenAndServe(port, router)

	if err != nil {
		log.Fatalf("Server cannot be started %v", err)
	}
}
