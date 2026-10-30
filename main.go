package main

import (
	"log"
	"net/http"

	"example.com/product-service/routes"
)

func main() {
	mux := http.NewServeMux()

	routes.SetupRoutes(mux)

	log.Println("Product Service http://localhost:8080 ünvanında başlayır")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Server xətası:", err)
	}
}