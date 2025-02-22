package elon

import "fmt"

// The Car that updates the number of meters driven based on the car's speed,
// and reduces the battery according to the battery drainage
func (car *Car) Drive() {
    // If batteryDrain is bigger than battery, battery is not decreased.
    if car.battery < car.batteryDrain {
        return
    }
    car.battery -= car.batteryDrain
    car.distance += car.speed
}

// To return the distance as displayed on the LED display as a string
func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// To return the battery percentage as displayed on the LED display as a string
func (car Car) DisplayBattery() string {
    return fmt.Sprintf("Battery at %d%%", car.battery)
}

// trackDistance int as its parameter and returns true if the car can finish the race;
// otherwise, return false:
func (car Car) CanFinish(trackDistance int) bool {
    batteryCount := car.battery / car.batteryDrain
    if trackDistance <= batteryCount * car.speed {
        return true
    }
    return false
}





package elon

// Car implements a remote controlled car.
type Car struct {
	speed        int
	batteryDrain int

	battery  int
	distance int
}

// NewCar creates a new car with given specifications.
func NewCar(speed, batteryDrain int) *Car {
	return &Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
	}
}