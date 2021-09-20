package main

import "fmt"

func main() {
	// if we use this greet ,_ := greeting("string")
	// _ means ignore the returned variable
	greet , num := greeting("Ahmed Hesham")
	fmt.Println(greet ,num)
	ConstructUser()
	ConstructUser2()
}

/*
* args => argName argDataType
* for 1 return
* funName(args) returnDataType {}
* for multiple return (please return by same orders)
* funName(args) (return_1_DT ,return_2_DT)
*/
func greeting(name string) (string ,int) {
	return  "Greeting, " + name, 100
}

/**
function explain 
case 1 : func name(arg) {} ,no return type VOID
case 2 : func name(arg) returnDT {} ,return single type
case 3 : func name(arg) (returnDT ,returnDT2) {} ,return multiple type [return type1 ,type2 ,....]
case 4 : func (struct) name(args) {} ,for using as a struct method
**/