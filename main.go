package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"goRest/databse"
	"goRest/routes"
	"log"
	"net/http"
)

func main() {
	// Load Configurations from config.json using Viper
	LoadAppConfig()

	// Initialize Database
	databse.Connect(AppConfig.ConnectionString)
	defer databse.Close()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)
	// Register Routes
	routes.RegisterEmployeeRoutes(router)
	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}
