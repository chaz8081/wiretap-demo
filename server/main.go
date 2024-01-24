package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var breadStore = BreadStore{}

func main() {
	err := breadStore.LoadData("data.json")
	if err != nil {
		log.Fatal("Failed to load data:", err)
	}

	http.HandleFunc("/api/breads", breadsHandler)

	fmt.Println("Server starting on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func breadsHandler(w http.ResponseWriter, r *http.Request) {
	breads := breadStore.GetAll()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(breads)
}
