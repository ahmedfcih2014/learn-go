package main

import "fmt"

func main() {
	// map is a key value datastructure
	// make(map[key data type] value data type)
	emails := make(map[int]string)

	emails[0] = "ahmedhesham2015fcih@gmail.com"
	emails[1] = "ahmedhesham2015fcih2@gmail.com"
	// emails["a"] = "ahmedhesham2015fcih2@gmail.com2" ,throw error

	fmt.Println(emails)
	fmt.Println(emails[1])
	// delete an element from the map by its key
	delete(emails ,1)
	fmt.Println(emails[1])

	// short hand
	emails2 := map[string]string{"ahmed":"ahmedhesham2015fcih@gmail.com" ,"gehan":"gehan9808@gmail.com"}

	fmt.Println(emails2)
	delete(emails2 ,"ahmed")
	fmt.Println(emails2)
}