package car

import (
	"fmt"
	"driver/vehicle"
	"driver/errors"
)

type Car struct {
	maxSpeed int
}

func New() vehicle.VehicleInterface {
	return &Car{
		maxSpeed: 200,
	}
}

func (car *Car) MoveFoward(speed int) (bool, error) {
	if (speed > car.maxSpeed) {
		return false, errors.New("Your car is too fast")
	}

	fmt.Println("Car is moving foward", speed)
	return true, nil
} 

