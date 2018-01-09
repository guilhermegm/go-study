package motorist

import (
	"fmt"
	"driver/vehicle"
)

func Drive(vehicle vehicle.VehicleInterface, speed int) {
	vehicle.MoveFoward(speed)
	fmt.Println("Motorist is driving")
}
