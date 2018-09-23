package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Serve(port string) {
	router := mux.NewRouter()
	router.HandleFunc("/getBike/{bikeID}", getBikeHandler)
	router.HandleFunc("/getAllBikes", getAllBikesHandler)
	router.HandleFunc("/addBike", addBikeHandler)

	log.Fatal(http.ListenAndServe(":"+port, router))
}

func getBikeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bikeID := vars["bikeID"]

	fmt.Fprintf(w, "Received a GET request for bike with ID %s.", bikeID)
}

func getAllBikesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Received a GET request for all bikes.")
}

func addBikeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Received a POST request to add a new bike.")
}
