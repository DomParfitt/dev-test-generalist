package main

import (
	"fmt"
	"github.com/DomParfitt/dev-test-generalist/src/dal/bike"
	"github.com/DomParfitt/dev-test-generalist/src/server"
	"os"
)

func main() {
	args := os.Args[1:]

	//If we don't have all the required args then exit
	if len(args) < 4 {
		fmt.Printf("Not enough arguments. Expected 4 but only %d provided.\n", len(args))
		printUsage()
		return
	}

	//Get the args into named vars
	url := args[0]
	db := args[1]
	collection := args[2]
	port := args[3]

	//Create the DAL
	bikeAccessor, err := bike.New(url, db, collection)

	//If we can't connect to URL, DB, Collection combo then exit
	if err != nil {
		fmt.Printf("Could not connect to MongoDB instance at URL: %s. Reason: %s.\n", url, err.Error())
		return
	}
	//Defer closing of connection
	defer bikeAccessor.Close()

	//Create and start API server
	server := server.New(bikeAccessor)
	server.Serve(port)
}

//Print the usage of this utility
func printUsage() {
	fmt.Println("Usage: main <url> <db_name> <collection_name> <port>")
	fmt.Println("\turl             - URL of the DB instance")
	fmt.Println("\tdb_name         - Name of the DB")
	fmt.Println("\tcollection_name - Name of the collection in the given DB")
	fmt.Println("\tport            - Port to serve the API on")

}
