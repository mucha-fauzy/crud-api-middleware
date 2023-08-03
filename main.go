package main

import (
	"crud_mysql_api_auth/infras"
	"crud_mysql_api_auth/internal/repository"
	"crud_mysql_api_auth/internal/services"
	"crud_mysql_api_auth/transport/middleware"
	"crud_mysql_api_auth/transport/routes"
	"fmt"
	"net/http"
)

func main() {
	// Create a new database connection
	db := infras.ProvideConn()

	// Initialize the repository with the database connection
	repo := repository.ProvideRepo(&db)

	// Initialize the service with the repository
	svc := services.ProvideService(repo)

	// Initialize the authentication middleware
	auth := middleware.ProvideAuthentication(&db)

	// Initialize the router with the service and authentication
	r := routes.ProvideRouter(svc, auth)

	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", r.SetupRoutes())
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}
