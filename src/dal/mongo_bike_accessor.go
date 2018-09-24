package dal

import (
	"github.com/DomParfitt/dev-test-generalist/src/common"
	"github.com/globalsign/mgo"
)

type MongoBikeAccessor struct {
	session *mgo.Session
}

func New(url string) (*MongoBikeAccessor, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}

	return &MongoBikeAccessor{session: session}, nil
}

func (m *MongoBikeAccessor) getBike(id int) (common.Bike, error) {

}

func (m *MongoBikeAccessor) getAllBikes() []common.Bike {

}

func (m *MongoBikeAccessor) addNewBike(bike common.Bike) {

}
