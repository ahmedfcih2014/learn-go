package main

import "fmt"

// whole user impl.
// fields
type User struct {
	FirstName string
	LastName string
	FullName string
	Age int
	AfterSomeYears int
	GetSixten int
	SomeYears int
}
// fields

// methods
func (user User) ComputFullName() User {
	user.FullName = user.FirstName + " " + user.LastName
	return user
}

func (user User) ComputAfterSomeYears() User {
	user.AfterSomeYears = user.Age + user.SomeYears
	return user
}

func (user User) ComputWhenGetSixten() User {
	user.GetSixten = 60 - user.Age
	return user
}

func (user User) PrintMyData() {
	fmt.Println("Im " ,user.FullName)
	fmt.Println("I'll get 60 after " ,user.GetSixten)
	fmt.Println("I'll be " ,user.AfterSomeYears ," After" ,user.SomeYears)
}
// methods

// act as constructor
func c() func(user User) User {
	return func(user User) User {
		return user
	}
}
//end of impl.

func main() {
	userConstructor := c()
	u := userConstructor(User{
		FirstName: "Ahmed",
		LastName: "Hesham",
		Age: 26,
		SomeYears: 15,
	}).
	ComputFullName().
	ComputAfterSomeYears().
	ComputWhenGetSixten()
	
	u.PrintMyData()
}