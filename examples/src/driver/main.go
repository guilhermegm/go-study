package main

import (
	"fmt"
	"driver/car"
	"driver/bus"
	"driver/motorist"
)

func main() {
	myCar := car.New()
	myBus := bus.New()

	fmt.Println("Starting...")

	motorist.Drive(myCar, 100)
	motorist.Drive(myBus, 50)
	motorist.Drive(myCar, 500)
}
