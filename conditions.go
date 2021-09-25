package main

import "fmt"

func main() {
	colors := []string{"Red" ,"Blue" ,"Orange" ,"Green"}
	
	colorsLen := len(colors)

	for i := 0 ;i < colorsLen ;i++ {
		color := colors[i]
		if color == "Red" {
			fmt.Println("Color is Red")
		} else if color == "Blue" {
			fmt.Println("Color is Blue")
		} else {
			fmt.Println("Color is not Red nor Blue")
		}
	}

	switch colors[1] {
	case "Red":
		fmt.Println("Color is Red")
	case "Blue":
		fmt.Println("Color is Blue")
	default:
		fmt.Println("Color is not Red nor Blue")
	}

	// FizzBuzz

	for i := 0; i <= 100; i++ {
		if i % 15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}