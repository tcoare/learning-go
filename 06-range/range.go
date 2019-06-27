package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5}
	sum := 0

	//range is used to iterate over elements and return relevant data
	//range returns two values, the index and the value

	//here we only care about the value so we use a blank identifier for the index.
	//we need to use a blank identifier because the compiler wont allow a variable
	//that is declared but not used.
	for _, nums := range numbers {
		sum += nums
	}
	fmt.Println(sum)

	//but there may be cases where we need to know the index of an element
	for pos, nums := range numbers {
		if nums == 4 {
			fmt.Println(pos)
		}
	}

	nameLetters := map[string]int{"Tyler": 5, "John": 4}

	//range is used on a map to iterate through its key value pairs
	for key, value := range nameLetters {
		fmt.Println(key, value)
	}

	//range can also be used to iterate just through the keys
	for keys := range nameLetters {
		fmt.Println(keys)
	}

}
