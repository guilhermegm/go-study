package motorist

import (
	"fmt"
	"driver/vehicle"
)

func Drive(vehicle vehicle.VehicleInterface, speed int) {
	_, err := vehicle.MoveFoward(speed)
	if (err != nil) {
		fmt.Println(err)
		return
	}
	fmt.Println("Motorist is driving")
}
