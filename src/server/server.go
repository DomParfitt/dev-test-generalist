package server

import (
	"encoding/json"
	"fmt"
	"github.com/DomParfitt/dev-test-generalist/src/common"
	"github.com/DomParfitt/dev-test-generalist/src/dal/bike"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

//Server structure
type Server struct {
	bikeAccessor bike.Accessor
}

//New server with a given Bike Accessor
func New(bikeAccessor bike.Accessor) *Server {
	return &Server{bikeAccessor: bikeAccessor}
}

//Serve the API on the given port
func (s *Server) Serve(port string) {
	router := mux.NewRouter()
	router.HandleFunc("/getBike/{bikeID}", s.getBikeHandler)
	router.HandleFunc("/getAllBikes", s.getAllBikesHandler)
	router.HandleFunc("/addBike", s.addBikeHandler)

	common.Log(fmt.Sprintf("Listening on port %s", port))
	log.Fatal(http.ListenAndServe(":"+port, router))
}

//Handler for getting a single bike by its ID
func (s *Server) getBikeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	//Attempt conversion of the bikeID to an int
	bikeID, err := strconv.Atoi(vars["bikeID"])
	if err != nil {
		common.Log(err.Error())
		writeErrorResponse(w, fmt.Sprintf("Please provide a numeric ID for a bike"))
		return
	}

	common.Log(fmt.Sprintf("Received a GET request for bike with ID %d.", bikeID))

	//Get the bike using the DAL
	bike, err := s.bikeAccessor.GetBike(bikeID)

	//Couldn't find a bike with that ID so return an error response
	if err != nil {
		common.Log(err.Error())
		writeErrorResponse(w, fmt.Sprintf("Could not retrieve bike with ID %d", bikeID))
		return
	}

	writeSuccessResponse(w, []common.Bike{bike})
}

//Handler for getting all bikes
func (s *Server) getAllBikesHandler(w http.ResponseWriter, r *http.Request) {
	common.Log("Received a GET request for all bikes.")

	//Get all bikes using DAL
	bikes := s.bikeAccessor.GetAllBikes()

	writeSuccessResponse(w, bikes)

}

//Handler for adding a new bike to the collection
func (s *Server) addBikeHandler(w http.ResponseWriter, r *http.Request) {

	//Get the POST data and write it into a bike struct
	decoder := json.NewDecoder(r.Body)
	bike := common.Bike{}
	err := decoder.Decode(&bike)

	//Post data is invalid, write to log and return error
	if err != nil {
		common.Log(err.Error())
		writeErrorResponse(w, err.Error())
		return
	}

	common.Log(fmt.Sprintf("Received a POST request to add a new bike:\n\t%v", bike))

	//Require bike to have at least a name, if not return error
	if bike.Name == "" {
		common.Log("Received request to add bike without a name.")
		writeErrorResponse(w, "Cannot add a bike without a name.")
		return
	}

	//Add the bike using the DAL
	err = s.bikeAccessor.AddBike(bike)

	//If we couldn't add the bike log the error and return it
	if err != nil {
		common.Log(err.Error())
		writeErrorResponse(w, err.Error())
		return
	}

	writeSuccessResponse(w, []common.Bike{bike})

}

//Write a successful response back to the originator including any relevant
//bike data
func writeSuccessResponse(w http.ResponseWriter, bikes []common.Bike) {
	response := Response{true, bikes, ""}
	writeResponse(w, response)
}

//Write an error response back to the originator with the error message set
func writeErrorResponse(w http.ResponseWriter, errorMsg string) {
	response := Response{false, []common.Bike{}, errorMsg}
	writeResponse(w, response)
}

//Write a response back to the originator
func writeResponse(w http.ResponseWriter, response Response) {
	json, err := json.Marshal(&response)

	if err != nil {
		fmt.Fprint(w, "There was an error processing the request.")
	}

	common.Log(fmt.Sprintf("Returning response %s", json))
	fmt.Fprintf(w, "%s", json)
}

//Response structure for replying to requests
type Response struct {
	Success  bool
	Bikes    []common.Bike
	ErrorMsg string
}
