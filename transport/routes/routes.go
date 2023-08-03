package routes

import (
	"crud_mysql_api_auth/internal/services"
	custom_middleware "crud_mysql_api_auth/transport/middleware"
	"net/http"

	"crud_mysql_api_auth/internal/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Router struct {
	Handler        *handlers.Handler
	Authentication *custom_middleware.Authentication
}

func ProvideRouter(service services.Service, auth *custom_middleware.Authentication) *Router {
	handler := handlers.ProvideHandler(service)
	return &Router{
		Handler:        handler,
		Authentication: auth,
	}
}

func (r *Router) SetupRoutes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(r.Authentication.GenerateRequestID)

	// Middleware for authentication using api key
	mux.Use(r.Authentication.SetXApiKey)
	mux.Use(r.Authentication.XApiAuthentication)

	mux.Route("/v1", func(rc chi.Router) {
		rc.Post("/products", r.Handler.CreateProduct)
		rc.Get("/products", r.Handler.ListProducts)
		rc.Put("/variants/{variantID}", r.Handler.UpdateVariant)
		rc.Delete("/products/{productID}", r.Handler.SoftDeleteProduct)
		rc.Delete("/products/hard/{productID}", r.Handler.HardDeleteProduct)
	})
	return mux
}
