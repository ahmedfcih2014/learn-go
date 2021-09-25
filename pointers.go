package main

import "fmt"

func main() {
	a := 5
	b := &a

	// show memory address
	fmt.Println(a ,b)

	// show address value
	fmt.Println(a ,*b)

	// here we have changed the address value which is reflect in (a variable)
	*b = 10
	fmt.Println(a)
	add := Add()
	fmt.Println(add(5 ,10))
}

func Add() func (num1 ,num2 int) int {
	return func (num1 ,num2 int) int {
		return num1 + num2
	}
}