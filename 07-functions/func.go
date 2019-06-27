package main

import (
	"errors"
	"fmt"
)

func main() {

	//functions are called using functionName(arguments)
	result := addTwo(3, 2)
	fmt.Println(result)

	//here the function is being assigned to multiple values
	//a result variable (result2) and an error variable (err)
	result2, err := divide(4, 0)
	if err != nil {
		fmt.Println(err) //Divide by 0 error
		//if an error is returned then err is not nil so print the error
	} else {
		fmt.Println(result2)
		//while there is no error print the result from the function
	}

	result3, err := divide(4, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result3)
	}

	//the variadic function is being called here and taking multiple arguments
	variadic := sum(1, 2, 3, 4, 5)
	fmt.Println(variadic)
}

//function that takes 2 integers and returns the sum as an integer
func addTwo(x int, y int) int {
	return x + y
}

//consecutive arguments of the same type can have their type declared once at the end
//note that this function returns 2 values, an integer and an error
func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Divide by 0 error")
	}
	return x / y, nil
}

//Go supports variadic functions which can take multiple values as an argument
//using '...' before the type indicates that the function will take more than one value
func sum(a ...int) int {
	total := 0
	for _, num := range a {
		total += num
	}
	return total
}
