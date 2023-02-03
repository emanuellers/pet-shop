package app

import (
	"log"
	"net/http"
	"pet_shop_app/handlers"

	"github.com/gorilla/mux"
)

func Init() {
	router := mux.NewRouter()

	//Routes
	router.HandleFunc("/pets", handlers.GetAllPets).Methods(http.MethodGet)
	router.HandleFunc("/pets/{petId:[0-9]+}", handlers.GetPetById).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:3000", router))
}
