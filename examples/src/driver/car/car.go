package car

import (
	"fmt"
	"driver/vehicle"
)

type Car struct {
	maxSpeed int
}

func New() vehicle.VehicleInterface {
	return &Car{
		maxSpeed: 200,
	}
}

func (car *Car) MoveFoward(speed int) {
	fmt.Println("Car is moving foward", speed)
} 

