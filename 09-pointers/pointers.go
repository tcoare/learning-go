package main

import "fmt"

func main() {

	num := 20

	//'&' is used to give the memory address of the variable
	fmt.Println(&num)

	//we can use '&' to store the address in another variable
	numPtr := &num
	fmt.Println(numPtr)

	//this is calling the function and passing in the value of num
	incrementValue(num)
	fmt.Println(num) //20

	//we call the function using '&' to indicate the pointer to the address of num
	incrementRef(&num)
	fmt.Println(num) //21

}

//the function here will have it arguments passed to it by value
//we are only receiving a copy of num in this function
//because it does not return any value the incremented number is discarded
func incrementValue(x int) {
	x++
}

//while this function uses the '*int' type meaning it takes the pointer as the value
func incrementRef(x *int) {

	//using '*' dereferences the the pointer from its memory address
	//without '*' we increment the memory address, not the value it holds
	//assigning a value to a dereferenced pointer changes it value
	*x++
}
