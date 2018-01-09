package bus

import (
	"fmt"
	"driver/vehicle"
)

type Bus struct {
	maxSpeed int
}

func New() vehicle.VehicleInterface {
	return &Bus{
		maxSpeed: 200,
	}
}

func (bus *Bus) MoveFoward(speed int) {
	fmt.Println("Bus is moving foward", speed)
} 

