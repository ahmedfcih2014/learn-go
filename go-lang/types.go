package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numInt int
	var numFloat float32
	var numUInt uint
	var boolV bool

	numInt = -32// - & +
	numFloat = 32.3
	numUInt = 32
	boolV = true

	fmt.Println(numInt)
	fmt.Println(numUInt)
	fmt.Println(numFloat)
	fmt.Println(boolV)

	age := 26
	name := "Ahmed Hesham"
	// strconv is a built in package to deal with strings
	fmt.Println(name + " " + strconv.Itoa(age) + " years old")
}