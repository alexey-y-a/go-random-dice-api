package main

import (
	"encoding/json"
	"math/rand/v2"
	"net/http"
)

func rollHandler(w http.ResponseWriter, r *http.Request) {
	result := rand.IntN(6) + 1
	json.NewEncoder(w).Encode(map[string]int{
		"result": result})
}

func main() {
	http.HandleFunc("/roll", rollHandler)
	http.ListenAndServe(":8080", nil)
}
