package dal

import (
	"github.com/DomParfitt/dev-test-generalist/src/common"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

//MongoBikeAccesor is a concrete implementation of the
//BikeAccessor interface for accessing the data in a
//MongoDB database
type MongoBikeAccessor struct {
	session *mgo.Session
}

//New creates a new instance of the MongoBikeAccessor
//with a access to the MongoDB instance provided by the
//given URL
func New(url string) (*MongoBikeAccessor, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}

	return &MongoBikeAccessor{session: session}, nil
}

//GetBike with the given ID from the connected MongoDB database.
func (m *MongoBikeAccessor) GetBike(id int) (*common.Bike, error) {
	bike := common.Bike{}
	err := m.session.DB("").C("").Find(bson.M{"id": id}).One(&bike)

	if err != nil {
		return nil, err
	}

	return &bike, nil
}

//GetAllBikes from the connceted MongoDB database
func (m *MongoBikeAccessor) GetAllBikes() []common.Bike {
	bikes := []common.Bike{}
	return bikes
}

//AddNewBike to the collection of Bikes in the MongoDB instance
func (m *MongoBikeAccessor) AddNewBike(bike *common.Bike) {

}

//Close the connection to the MongoDB instance
func (m *MongoBikeAccessor) Close() {
	m.session.Close()
}
