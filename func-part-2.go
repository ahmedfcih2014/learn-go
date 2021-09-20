package main

import "fmt"

type User struct {
	Name string
	Email string
	Age int
}

// by value
// how to define method to struct
func (user User) changeName() User {
	user.Name = "Changed Name"
	return user
}

// by reference
func (user *User) changeRefName() {
	user.Name = "Changed Reference Name"
}

func ConstructUser() {
	// create and set data
	var user User
	user.Name = "Ahmed"
	user.Email = "ahmedhesham@asd.com"
	user.Age = 26

	// most take the return as variable and use it
	user.changeName()
	fmt.Println(user.Name)
	fmt.Println("His age:" ,user.Age)
	fmt.Println("Contact him by mail:" ,user.Email)
}

func ConstructUser2() {
	// create with initialized user
	user2 := User{
		Name: "Ali",
		Email: "alihesham@asd.com",
		Age: 15,
	}
	user2.changeRefName()
	fmt.Println(user2.Name)
	fmt.Println("His age:" ,user2.Age)
	fmt.Println("Contact him by mail:" ,user2.Email)
}