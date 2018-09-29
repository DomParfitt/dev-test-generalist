package bike

import (
	"github.com/DomParfitt/dev-test-generalist/src/common"
)

//Accessor defines an interface for types that can
//access a collection of Bike objects, e.g. in a database
type Accessor interface {
	GetBike(id int) (common.Bike, error)
	GetAllBikes() []common.Bike
	AddBike(bike common.Bike)
}
