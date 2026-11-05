package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"io"
	"example.com/product-service/models"
	"strings"
	"example.com/product-service/database"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
			getProducts(w, r)

	case http.MethodPost:
		createProduct(w, r)

	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(
			w,
			 "Yalniz GET ve POST metodlarina icaze verilir.",
			 http.StatusMethodNotAllowed,
		)
	}
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	products, err := database.GetProducts()
	if err != nil {
		log.Println("mehsullar oxunarken xeta:", err)
		http.Error(w,
			 "Mehsullar oxuna bilmedi", 
			 http.StatusInternalServerError,)
			 return
		}
		w.Header().Set("Content-Type", "application/json")

		err = json.NewEncoder(w).Encode(products)
		if err != nil {
			log.Println("Mehsullar gonderilerken xeta:", err)
		}
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name 		string 		`json:"name"`
		Price 		int64		`json:"price"`
		Stock 		int 		`json:"stock"`
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1024*1024) 

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&input)
	if err != nil {
		http.Error(
			w,
			"Duzgun JSON melumati gonder",
			http.StatusBadRequest,
		)
		return
	}	

	err = decoder.Decode(&struct{}{})
	if err != io.EOF {
		http.Error(
			w,
			"Yalniz bir JSON obyekt gonderin",
			http.StatusBadRequest,
		)
		return
	} 
 
	input.Name = strings.TrimSpace(input.Name)

	if input.Name == "" || input.Price <= 0 || input.Stock < 0 {
		http.Error(
			w,
			"Ad bos, qiymet ve stok menfi ola bilmez",
			http.StatusBadRequest,
		)
		return
	}

	product := models.Product{
		Name: input.Name,
		Price: input.Price,
		Stock: input.Stock,
	}

	created, err := database.CreateProduct(product)
	if err != nil {
		log.Println("Mehsul saxlanilarken xeta:", err)
		http.Error(
			w,
			"Mehsul saxlanilamadi",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(created)
	if err != nil {
		log.Println("Yeni Mehsul gonderilerken xeta:", err)
	}
}