package main

import (
	"fmt"
)

func main() {
	p1 := Product{"Vando"}
	p2 := Product{"Junim"}

	cart := NewCart()

	cart.addProduct(p1, 1)
	fmt.Println(cart)

	cart.addProduct(p2, 5)
	fmt.Println(cart)

	cart.setQuantity(p1, 17)
	fmt.Println(cart)

	cart.setQuantity(p1, 3)
	fmt.Println(cart)

	cart.removeProduct(p1)
	fmt.Println(cart)
}
