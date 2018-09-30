package common

//Bike stucture that matches the schema of the DB
type Bike struct {
	BikeID      int "bikeId"
	Name        string
	Description string
	Price       string
}
