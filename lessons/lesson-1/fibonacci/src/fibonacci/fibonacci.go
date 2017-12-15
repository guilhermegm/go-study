package main;

import "fmt"

func main() {
	a, b := 0, 1

	for i := 0; i < 50; i++ {
		a = a + b
		a, b = b, a

		fmt.Println(b)
	}
}
