package dal

import (
	"github.com/DomParfitt/dev-test-generalist/src/common"
)

//BikeAccessor defines an interface for types that can
//access a collection of Bike objects, e.g. in a database
type BikeAccessor interface {
	getBike(id int) (common.Bike, error)
	getAllBikes() []common.Bike
	addBike(bike common.Bike)
}
