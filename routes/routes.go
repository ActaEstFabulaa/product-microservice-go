package routes

import (
	"net/http"
	"example.com/product-service/handlers"
)

func SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/product", handlers.ProductHandler)
}