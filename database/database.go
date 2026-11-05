package database

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"example.com/product-service/models"
)

const productsFile = "data/products.json"

var fileMutex sync.Mutex

func GetProducts() ([]models.Product, error) {
	fileMutex.Lock()
	defer fileMutex.Unlock()
	
	return readProducts()

}

func CreateProduct(product models.Product) (models.Product, error){
	fileMutex.Lock()
	defer fileMutex.Unlock()

	products, err := readProducts()
	if err != nil {
		return models.Product{}, err
	}

	maxID := 0
	for _, existing := range products {
		if existing.ID > maxID {
			maxID = existing.ID
		}
	}
	product.ID = maxID +1 
	products = append(products, product)

	err = writeProducts(products)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil 
}

func readProducts() ([]models.Product, error) {
	data,err := os.ReadFile(filepath.Join(productsFile))
	if err != nil {
		return nil, err
	}

	products := make([]models.Product, 0)

	err = json.Unmarshal(data, &products)
	if err != nil {
		return nil, err
	}

    return products, nil
}

func writeProducts(products []models.Product) error {
	data, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		return err
	}

	directory := filepath.Dir(productsFile)
	
	
	tempFile, err := os.CreateTemp(directory, "products_*.json")
	if err != nil {
		return err
	}	
	
	tempPath := tempFile.Name()
	defer os.Remove(tempPath)

	_, err = tempFile.Write(data)
	if err != nil {
		tempFile.Close()
		return err
	}
	
	err = tempFile.Close()
	if err != nil {
		return err
	}	

	return os.Rename(tempPath, productsFile)
}