package main

import (
	"encoding/json"
	"net/http"
)

type Pet struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

func getAllPets(w http.ResponseWriter, r *http.Request) {
	pets := []Pet{
		{Name: "Lulu", Owner: "Dalila"},
		{Name: "Thor", Owner: "Pablo"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pets)
}
func main() {
	//Routes
	http.HandleFunc("/pets", getAllPets)

	http.ListenAndServe("localhost:3000", nil)
}
