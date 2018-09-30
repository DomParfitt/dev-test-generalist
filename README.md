# Dev Test

## Getting Started

### Requirements
1. Golang v1.11 or higher

### Installation
Two build files are provided in the root of the project for Linux (build.sh) and Windows (build.bat) respectively. To build the code simple run the relevant file. This will create the file `main.exe` and place it in the `bin/` directory at the root of the project.

### Running the Code
1. Install the binary from the above instructions.
2. Run the binary from the command line using the following command:
```
main <url> <db_name> <collection_name> <port>
```
Where:
 - `url` is the URL that the MongoDB instance is being served on.
 - `db_name` is the name of the DB on the MongoDB instance
 - `collection_name` is the name of the collection on the given DB holding the bike data
 - `port` is the port that you wish to serve the API on

### Database instructions

Follow these instructions to get the test database working on your machine:

1. Ensure you have the latest version of Docker installed on your machine (Native Docker for Windows, and Docker for Mac that no longer use docker-machine) [https://docs.docker.com/engine/installation/](https://docs.docker.com/engine/installation/).
2. Ensure the docker service is running on your machine and you can connect to it using the `docker info` command
3. Get Mongo running as a service on your machine by typing the following commands into a console window.
```bash
#A. remove the old instance of the db if it exists, don't worry if this errors 
docker rm --force jlmongo

#B. start the Mongo container as a service
docker run -d --name jlmongo -p 27017:27017 jujhars13/dev-test-generalist-mongo:latest

#C. Once the container is up and running. Import the bike schema by running this command in
docker exec jlmongo mongoimport --collection bike /schema/bike.json --jsonArray

```

## API
All calls to the API return a JSON Response object which conforms to the below schema:
```
Response {
  Success: boolean
  Bikes: Bike[]
  ErrorMsg: string
}
```

Successful requests return a `Response` with `Success` set to `true`, the `Bikes` array containing any relevant bike data and `ErrorMsg` set to the empty string.

Unsuccessful requests return a `Response` ith `Success` set to `false`, the `Bikes` array empty and `ErrorMsg` containing details of the error.


`Bikes` are defined as follows:
```
Bike {
  BikeID: int
  Name: string
  Description: string
  Price: string
}
```
A valid `Bike` object requires at least the `Name` field be present.

### Get Bike
```
/getBike/{bikeID}
```
`GET` request to retrieve a single bike by its ID. If there is a bike with the given ID a success `Response` is returned with a single bike in the `Bikes` array, which matches the provided ID.

#### Example
```
curl -s localhost:8080/getBike/1
```

### Get All Bikes
```
/getAllBikes
```
`GET` request to retrieve all bikes held in the collection. 

#### Example
```
curl -s localhost:8080/getAllBikes
```

### Add Bike
```
/addBike
```
`POST` request to add a new bike to the collection. The data for the `POST` request must be a valid `Bike` object, i.e. must have at least a `Name`.

If a `BikeID` is provided then the `Bike` will be added only if there is not already an existing `Bike` with the same ID. If no `BikeID` is provided then the `Bike` will be added with the next available `BikeID`.

#### Examples
##### Bike with just a Name
```
curl -s localhost:8080/addNewBike -d '{"Name": "New Bike"}'
```

##### Bike with ID and Name
```
curl -s localhost:8080/addNewBike -d '{"BikeID": 10, "name": "New Bike"}'
```

##### Bike with all fields
```
curl -s localhost:8080/addNewBike -d '{"BikeID": 10, "name": "New Bike", "Description": "A new bike" "Price": "1000"}'
```

## Assumptions
1. Security/authentication will be handled by another layer.
2. A `Bike` must have at least a `Name` to be considered valid
3. Attempting to add a `Bike` with a `BikeID` that is already used results in an error, rather than updating/overwriting the existing `Bike`