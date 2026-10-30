package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"example.com/product-service/models"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(
			w,
			"Yalnız GET metoduna icazə verilir",
			http.StatusMethodNotAllowed,
		)
		return
	}

	products := []models.Product{
		{
			ID:    1,
			Name:  "Klaviatura",
			Price: 2550,
			Stock: 10,
		},
		{
			ID:    2,
			Name:  "Siçan",
			Price: 1800,
			Stock: 20,
		},
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(products)
	if err != nil {
		log.Println("Məhsullar göndərilərkən xəta:", err)
	}
}