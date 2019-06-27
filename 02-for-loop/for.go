package main

import "fmt"

func main() {

	n := 1

	//Single condition for loop, equivalent of a while loop
	for n <= 3 {
		fmt.Println(n)
		n = n + 1
	}

	//initial; condition; after loop
	for i := 4; i < 6; i++ {
		fmt.Println(i)
	}

	//for statements with no condition will loop until you break out of them
	for {
		fmt.Println("looping")
		break
	}

	//continue can be used to move to the next iteration of the loop
	for num := 0; num <= 10; num++ {
		if num%2 == 0 {
			continue
			//if number is even will move to next iteration
		}
		fmt.Println(num)
	}
}
