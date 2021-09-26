package main

import "fmt"
func main() {
	sum := func (num1 int ,num2 int) int {return num1 + num2}
	str ,num := computed(sum)
	fmt.Println(str)
	fmt.Println(num)
	// add := add()
	fmt.Println(add()(200 ,20))
}

// func take func as an argument
func computed(sum func (num1 int ,num2 int) int) (string ,int) {
	v := sum(10 ,20)
	return "value is" ,v
}

// closure (return function)
func add() func(num1 int, num2 int) int {
	return func(num1 int, num2 int) int {
		return num1 + num2
	}
}