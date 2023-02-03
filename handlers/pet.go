package handlers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"pet_shop_app/schemas"

	"github.com/gorilla/mux"
)

func GetAllPets(w http.ResponseWriter, r *http.Request) {
	pets := []schemas.Pet{
		{Name: "Lulu", Owner: "Dalila"},
		{Name: "Thor", Owner: "Pablo"},
	}

	acceptedType := r.Header.Get("Accept")

	if acceptedType == "application/json" {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pets)
	} else if acceptedType == "application/xhtml+xml" {
		w.Header().Add("Content-Type", "application/xhtml+xml")
		xml.NewEncoder(w).Encode(pets)
	}

}

func GetPetById(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	fmt.Fprint(w, pathParams["petId"])
}
