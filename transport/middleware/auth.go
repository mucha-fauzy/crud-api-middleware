package middleware

import (
	"context"
	"crud_mysql_api_auth/infras"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Authentication struct {
	DB *infras.Conn
}

const (
	HeaderAuthorization = "X-Api-Key"
)

func ProvideAuthentication(db *infras.Conn) *Authentication {
	return &Authentication{
		DB: db,
	}
}

// For demonstration purposes, I hardcode a mock X-Api-Key
const mockApiKey = "this_is_a_mock_key"

// Middleware for authentication using api key
func (a *Authentication) XApiAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKeyHeader := r.Header.Get(HeaderAuthorization)
		apiKey := mockApiKey
		log.Println("Running XApiAuthentication")
		if apiKeyHeader == apiKey {
			log.Println("User unauthorized")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)

			return
		}
		next.ServeHTTP(w, r)
	})
}

// Middleware set api key
func (a *Authentication) SetXApiKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		apiKey := mockApiKey

		// Set the X-Api-Key header
		req.Header.Set(HeaderAuthorization, apiKey)

		// Call the next handler in the chain
		next.ServeHTTP(w, req)
	})
}

// Middleware to generate and include the requestId in the request context
func (a *Authentication) GenerateRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String() + "-" + "mock"
		ctx := context.WithValue(r.Context(), "requestId", requestID)
		r = r.WithContext(ctx)
		log.Println("Running GenerateRequestID")
		w.Header().Set("X-Request-ID", requestID)
		next.ServeHTTP(w, r)
	})
}
