package main

import (
	"fmt"
	"strconv"
)

func main() {
	// normal for loop
	// for i:= 1; i <= 5; i++ {
	// 	fmt.Println(i)
	// }


	// act like while
	// count := 1

	// for count <= 5 {
	// 	fmt.Println(count)
	// 	count++
	// }

	// don`t run this please
	// for {
	// 	fmt.Println("Infinity loop")
	// }

	for i:= 1; i <= 10; i++ {
		if i <= 5 {
			fmt.Println("Continue blook i: " ,strconv.Itoa(i))
			continue
		}
		fmt.Println(i)
		if i >= 7 {
			fmt.Println("break blook i: " ,strconv.Itoa(i))
			break
		}
	}
}