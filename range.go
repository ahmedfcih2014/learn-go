package main

import "fmt"

func main() {
	// range is pwoerful with key value DS
	ids := []int{3,5,10,12,6,35}

	for index ,value := range ids {
		fmt.Printf("index %d contains id : %d\n" ,index ,value)
	}

	emails := map[string]string{"ahmed":"ahmedhesham2015fcih@gmail.com" ,"gehan":"gehan9808@gmail.com"}

	for key ,value := range emails {
		fmt.Printf("%s: %s\n" ,key ,value)
	}
}