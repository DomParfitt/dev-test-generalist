package dal

import (
	"github.com/DomParfitt/dev-test-generalist/src/common"
)

//BikeAccessor defines an interface for types that can
//access a collection of Bike objects, e.g. in a database
type BikeAccessor interface {
	GetBike(id int) (common.Bike, error)
	GetAllBikes() []common.Bike
	AddBike(bike common.Bike)
}
