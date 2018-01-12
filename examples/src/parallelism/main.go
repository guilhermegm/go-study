package main

import (
	"fmt"
	"time"
)


func outraGoRoutine() {
	for {
		fmt.Println("Vandinho")
		time.Sleep(time.Second * 1)
	}
}

func sender(c chan string) {
	for {
		c <- "channel 1"
	}
}

func printer(c chan string) {
	for {
		msg := <- c
		fmt.Println(msg + " channel 2")
		time.Sleep(time.Second * 2)
	}
}

func main() {
	exampleChannel := make(chan string)

	go sender(exampleChannel)
	go printer(exampleChannel)
	go outraGoRoutine()

	var input string
	fmt.Scanln(&input)
}
