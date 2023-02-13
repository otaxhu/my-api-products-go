package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Esta es la raiz =)")
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	productsDB := "./prueba-03-api-products/productsdb.json"
	//file, err1 := os.Open("")
	/*
		if err1 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err1)
			return
		}
	*/
	//defer file.Close()
	bytes, err := os.ReadFile(productsDB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	var products []Products
	err = json.Unmarshal(bytes, &products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	for _, product := range products {
		fmt.Fprintf(w, "ID: %d\nNombre: %s\nPrecio: %d\n", product.ID, product.Name, product.Price)
	}
}

func PostProducts(w http.ResponseWriter, r *http.Request) {
	productsDB := "./prueba-03-api-products/productsdb.json"

	var product Products
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	product.ID, err = NextID()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
	}
	bytes, err := os.ReadFile(productsDB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	var products []Products

	err = json.Unmarshal(bytes, &products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	products = append(products, product)
	data, err := json.MarshalIndent(products, "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	err = os.WriteFile(productsDB, data, 0644)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func PutProductByID(w http.ResponseWriter, r *http.Request) {
	productsDB := "./prueba-03-api-products/productsdb.json"
	var updatedProduct Products
	err := json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	data, err := DecodeJSONfile(productsDB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	var updated bool
	for i, product := range data {
		if product.ID == updatedProduct.ID {
			if updatedProduct.Name == "" || updatedProduct.Price == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			data[i].Name = updatedProduct.Name
			data[i].Price = updatedProduct.Price
			updated = true
			break
		}
	}
	if !updated {
		w.WriteHeader(http.StatusNotFound)
	}
	bytes, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	err = os.WriteFile(productsDB, bytes, 0644)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
}

func DeleteProductByID(w http.ResponseWriter, r *http.Request) {
	productsDB := "./prueba-03-api-products/productsdb.json"
	var deletedProduct Products
	err := json.NewDecoder(r.Body).Decode(&deletedProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	data, err := DecodeJSONfile(productsDB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	var deleted bool
	for i, product := range data {
		if product.ID == deletedProduct.ID {
			data = append(data[:i], data[i+1:]...)
			deleted = true
			break
		}
	}
	if !deleted {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	bytes, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	err = os.WriteFile(productsDB, bytes, 0644)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
}
