package main

import "fmt"

func main() {
	// declare array set array length & its data types
	var fruits [2]string

	// assign values to array
	fruits[0] = "Apple"
	fruits[1] = "Orange"

	// declare & assign values
	fruits2 := [2]string{"Banana" ,"Pinapple"}

	fmt.Println(fruits)
	fmt.Println(fruits2)

	// another way is (Slice) its not need for set the length
	fruitSlice := []string{"Apple" ,"Banana" ,"Pinapple" ,"Orange" ,"Watermelon"}
	fmt.Println(fruitSlice)

	// startIndex:endBeforeIndex
	fmt.Println("part array" ,fruitSlice[1:4])
}