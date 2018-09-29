package server

import (
	"encoding/json"
	"fmt"
	"github.com/DomParfitt/dev-test-generalist/src/dal/bike"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	bikeAccessor bike.Accessor
}

//New server
func New(bikeAccessor bike.Accessor) *Server {
	return &Server{bikeAccessor: bikeAccessor}
}

//Serve the API on the given port
func (s *Server) Serve(port string) {
	router := mux.NewRouter()
	router.HandleFunc("/getBike/{bikeID}", s.getBikeHandler)
	router.HandleFunc("/getAllBikes", s.getAllBikesHandler)
	router.HandleFunc("/addBike", s.addBikeHandler)

	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

//Handler for getting a single bike by its ID
func (s *Server) getBikeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	bikeID, err := strconv.Atoi(vars["bikeID"])

	if err != nil {
		fmt.Fprint(w, "Please provide a numeric ID for a bike")
		return
	}

	fmt.Printf("Received a GET request for bike with ID %d.\n", bikeID)

	bike, err := s.bikeAccessor.GetBike(bikeID)

	if err != nil {
		fmt.Fprintf(w, "Could not retrieve bike with ID %d", bikeID)
		return
	}

	json, err := json.Marshal(bike)

	if err != nil {
		fmt.Fprintf(w, "Could not retrieve bike with ID %d", bikeID)
		return
	}

	fmt.Printf("Returning result %s\n", json)

	fmt.Fprintf(w, "%s", json)

}

//Handler for getting all bikes
func (s *Server) getAllBikesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received a GET request for all bikes.\n")

	bikes := s.bikeAccessor.GetAllBikes()

	json, err := json.Marshal(bikes)

	if err != nil {
		fmt.Fprintf(w, "Could not retrieve bikes")
		return
	}

	fmt.Printf("Returning result %s\n", json)

	fmt.Fprintf(w, "%s", json)
}

//Handler for adding a new bike to the collection
func (s *Server) addBikeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Received a POST request to add a new bike.")
}
