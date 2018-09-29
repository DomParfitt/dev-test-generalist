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
	session *mgo.Session
}

//New creates a new instance of the MongoBikeAccessor
//with a access to the MongoDB instance provided by the
//given URL
func New(url string) (MongoBikeAccessor, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return MongoBikeAccessor{session: nil}, err
	}

	return MongoBikeAccessor{session: session}, nil
}

//GetBike with the given ID from the connected MongoDB database.
func (m MongoBikeAccessor) GetBike(id int) (common.Bike, error) {
	bike := common.Bike{}
	err := m.session.DB("test").C("bike").Find(bson.M{"bikeId": id}).One(&bike)
	fmt.Printf("%v\n", bike)

	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return bike, err
	}

	return bike, nil
}

//GetAllBikes from the connceted MongoDB database
func (m MongoBikeAccessor) GetAllBikes() []common.Bike {
	bikes := []common.Bike{}
	err := m.session.DB("test").C("bike").Find(bson.M{}).All(&bikes)

	if err != nil {
		return []common.Bike{}
	}

	return bikes
}

//AddBike to the collection of Bikes in the MongoDB instance
func (m MongoBikeAccessor) AddBike(bike common.Bike) {

}

//Close the connection to the MongoDB instance
func (m *MongoBikeAccessor) Close() {
	m.session.Close()
}
