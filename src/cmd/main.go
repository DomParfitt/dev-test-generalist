package main

import (
	"fmt"
	"github.com/DomParfitt/dev-test-generalist/src/dal/bike"
	"github.com/DomParfitt/dev-test-generalist/src/server"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Unexpected number of arguments. Expected 1 - MongoDB URL")
		return
	}
	url := args[0]
	port := args[1]

	bikeAccessor, err := bike.New(url)
	if err != nil {
		fmt.Printf("Could not connect to MongoDB instance at URL: %s", url)
	}
	defer bikeAccessor.Close()

	server := server.New(bikeAccessor)
	server.Serve(port)
}

func printUsage() {
	fmt.Println("Usage")
	fmt.Println("\tURL - URL of the DB instance")
	fmt.Println("\tPort - Port to serve the API on")

}
