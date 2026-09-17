//An interface that delegates the instantiation of an object to its implementing structs, allowing behavior to vary while standardizing the creation signature.
package factory

import "fmt"

// 1. The product interface
type Vehicle interface {
	Drive()
}

type Car struct{}
func (c Car) Drive() { fmt.Println("Driving a car") }

type Bike struct{}
func (b Bike) Drive() { fmt.Println("Riding a bike") }

// 2. THE FACTORY METHOD (The Creator Interface)
type VehicleFactory interface {
	Create() Vehicle
}

// 3. Specific factories for each type
type CarFactory struct{}
func (cf CarFactory) Create() Vehicle {
	return Car{}
}

type BikeFactory struct{}
func (bf BikeFactory) Create() Vehicle {
	return Bike{}
}

// Usage:
// var factory VehicleFactory = CarFactory{}
// myRide := factory.Create()