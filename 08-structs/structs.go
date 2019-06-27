package main

import "fmt"

//this creates a struct type named person
//it has fields of name and age which have a type declared
type person struct {
	name string
	age  int
}

func main() {

	//this is the syntax used to create a srtuct
	firstPerson := person{name: "Tyler", age: 21}
	fmt.Println(firstPerson) //{Tyler 21}

	//structs can be initialised with the fields of the struct omitted
	secondPerson := person{"John", 30}
	fmt.Println(secondPerson) //{John 30}

	//values of a struct can be accessed with a dot
	fmt.Println(firstPerson.name) //Tyler
	fmt.Println(secondPerson.age) //30

	//using '&' before the struct is a pointer to the sruct
	firstPersonPtr := &firstPerson
	fmt.Println(firstPersonPtr) //&{Tyler 21}

	//Structs are mutable
	firstPerson.age = 22
	fmt.Println(firstPerson.age) //22

}
