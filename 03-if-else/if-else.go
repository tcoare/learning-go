package main

import "fmt"

func main() {

	//basic if else statement
	if 3%2 == 0 {
		fmt.Println("3 is even")
	} else {
		fmt.Println("3 is odd")
	}

	//An if statement can be used without an else statement
	if 10%5 == 0 {
		fmt.Println("10 is divisible by 5")
	}

	//A variable can be declared before the conditionals
	//Any variables declared in this statement can be used in all branches
	if num := 11; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
