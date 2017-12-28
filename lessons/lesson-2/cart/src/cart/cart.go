package main

import (
	"fmt"
)

type CartInterface interface {  
	addProduct() []Product
	removeProduct() []Product
	setQuantity() []Product
}

type Cart struct {
	products map[Product]int
}

func NewCart() Cart {
	cart := Cart{make(map[Product]int)}
	return cart
}

func (cart Cart) addProduct(product Product, quantity int) {
	cart.products[product] = quantity
}

func (cart Cart) removeProduct(product Product) {
	delete(cart.products, product)
}

func (cart Cart) setQuantity(product Product, quantity int) {
	cart.products[product] = quantity
}

func Test() {
	fmt.Println("123123")
}
