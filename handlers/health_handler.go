package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(
			w,
			"Yalnız GET metoduna icazə verilir",
			http.StatusMethodNotAllowed,
		)
		return
	}

	response := struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("Health cavabı göndərilərkən xəta:", err)
	}
}