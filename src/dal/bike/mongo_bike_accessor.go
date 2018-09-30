package bike

import (
	"fmt"
	"github.com/DomParfitt/dev-test-generalist/src/common"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

//MongoBikeAccessor is a concrete implementation of the
//BikeAccessor interface for accessing the data in a
//MongoDB database
type MongoBikeAccessor struct {
	session    *mgo.Session
	database   *mgo.Database
	collection *mgo.Collection
}

//New creates a new instance of the MongoBikeAccessor
//with a access to the MongoDB instance provided by the
//given URL
func New(url, databaseName, collectionName string) (MongoBikeAccessor, error) {

	//Attempt to open a session on the provided URL
	session, err := mgo.Dial(url)

	//Can't open a session, return error
	if err != nil {
		return MongoBikeAccessor{}, err
	}

	//Check the provided DB exists and if not return error
	databaseNames, err := session.DatabaseNames()
	if err != nil {
		return MongoBikeAccessor{}, err
	}

	if !contains(databaseName, databaseNames) {
		return MongoBikeAccessor{}, fmt.Errorf("no database with name %s", databaseName)
	}

	database := session.DB(databaseName)

	//Check the provided collection exists and if not return error
	collectionNames, err := database.CollectionNames()
	if err != nil {
		return MongoBikeAccessor{}, err
	}

	if !contains(collectionName, collectionNames) {
		return MongoBikeAccessor{}, fmt.Errorf("no collection with name %s", collectionName)
	}

	collection := database.C(collectionName)

	//Return the accessor
	return MongoBikeAccessor{session, database, collection}, nil
}

//GetBike with the given ID from the connected MongoDB database.
func (m MongoBikeAccessor) GetBike(id int) (common.Bike, error) {
	common.Log(fmt.Sprintf("Attempting to retrieve bike with ID %d", id))

	bike := common.Bike{}
	err := m.collection.Find(bson.M{"bikeId": id}).One(&bike)

	if err != nil {
		common.Log(err.Error())
		return bike, err
	}

	common.Log(fmt.Sprintf("Retrieved: %v", bike))
	return bike, nil
}

//GetAllBikes from the connceted MongoDB database
func (m MongoBikeAccessor) GetAllBikes() []common.Bike {
	common.Log("Attempting to retrieve all bikes.")

	bikes := []common.Bike{}
	err := m.collection.Find(bson.M{}).All(&bikes)
	if err != nil {
		common.Log(err.Error())
		return []common.Bike{}
	}

	common.Log(fmt.Sprintf("Retrieved %d bikes.", len(bikes)))
	return bikes
}

//AddBike to the collection of Bikes in the MongoDB instance
func (m MongoBikeAccessor) AddBike(bike common.Bike) error {
	common.Log(fmt.Sprintf("Attempting to add new bike: %v", bike))

	_, err := m.GetBike(bike.BikeID)
	if err == nil {
		common.Log(fmt.Sprintf("Attempted to add bike with existing ID: %d", bike.BikeID))
		return fmt.Errorf("Cannot add new bike as there is an existing bike with ID %d", bike.BikeID)
	}

	err = m.collection.Insert(bike)

	if err != nil {
		common.Log(err.Error())
		return err
	}

	common.Log("Successfully added a new bike.")
	return nil
}

//Close the connection to the MongoDB instance
func (m *MongoBikeAccessor) Close() {
	common.Log("Closing the session.")
	m.session.Close()
}

func contains(needle string, haystack []string) bool {
	for _, value := range haystack {
		if value == needle {
			return true
		}
	}

	return false
}
