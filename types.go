package main

import (
	"encoding/json"
	"os"
)

type Products struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func DecodeJSONfile(path string) ([]Products, error) {
	var data []Products
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func NextID() (int, error) {
	productsDB := "./prueba-03-api-products/productsdb.json"
	data, err := DecodeJSONfile(productsDB)
	if err != nil {
		return 0, err
	}
	var maxID int
	for _, product := range data {
		if product.ID > maxID {
			maxID = product.ID
		}
	}
	return maxID + 1, nil
}
