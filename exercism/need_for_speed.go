package speed

// define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
	}
}

// define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	// if battery is less than batteryDrain, the car can't move.
	if car.battery < car.batteryDrain {
		return car
	}
	car.battery -= car.batteryDrain
	car.distance += car.speed
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	carDriveCount := car.battery / car.batteryDrain
	// car.speed * carDriveCount == how long car move
	if track.distance <= car.speed*carDriveCount {
		return true
	}
	return false
}
