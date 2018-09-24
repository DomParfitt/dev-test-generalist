package server

import (
	"encoding/json"
	"fmt"
	"github.com/DomParfitt/dev-test-generalist/src/common"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//Serve the API on the given port
func Serve(port string) {
	router := mux.NewRouter()
	router.HandleFunc("/getBike/{bikeID}", getBikeHandler)
	router.HandleFunc("/getAllBikes", getAllBikesHandler)
	router.HandleFunc("/addBike", addBikeHandler)

	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

//Handler for getting a single bike by its ID
func getBikeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bikeID := vars["bikeID"]

	fmt.Printf("Received a GET request for bike with ID %s.\n", bikeID)

	bike := &common.Bike{
		BikeID:      1,
		Name:        "Dummy Bike",
		Description: "Placeholder data for a non-existant bike",
		Price:       "1,000,000",
	}

	json, err := json.Marshal(bike)

	if err != nil {
		fmt.Fprintf(w, "Could not retrieve bike with ID %s", bikeID)
		return
	}

	fmt.Printf("Returning result %s\n", json)

	fmt.Fprintf(w, "%s", json)

}

//Handler for getting all bikes
func getAllBikesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received a GET request for all bikes.\n")

	bikes := []common.Bike{}

	for i := 0; i < 5; i++ {
		bike := &common.Bike{
			BikeID:      i,
			Name:        "Dummy Bike",
			Description: "Placeholder data for a non-existant bike",
			Price:       "1,000,000",
		}

		bikes = append(bikes, *bike)
	}

	json, err := json.Marshal(bikes)

	if err != nil {
		fmt.Fprintf(w, "Could not retrieve bikes")
		return
	}

	fmt.Printf("Returning result %s\n", json)

	fmt.Fprintf(w, "%s", json)
}

//Handler for adding a new bike to the collection
func addBikeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Received a POST request to add a new bike.")
}
