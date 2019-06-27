package main

import "fmt"

func main() {

	//Create an array of a fixed size, the size of which is part of its type
	//By default an array is 0 vauled
	var a [5]int
	var b [5]string
	var c [5]bool

	fmt.Println(a, b, c) //[0 0 0 0 0] [    ] [false false false false false]

	//Value can be set using the index of the array
	a[1] = 20

	fmt.Println(a) //[0 20 0 0 0]

	//len builtin returns length of an array
	fmt.Println(len(a)) //5

	//An arrays contents can be initialised when it is declared
	d := [5]int{1, 2, 3, 4, 5}

	fmt.Println(d) //[1 2 3 4 5]

	//multi-dimensional arrays can be created when declaring as part of its type
	var multiDimensional [2][3]int

	fmt.Println(multiDimensional) //[[0 0 0] [0 0 0]]

	//slices can be created using the builtin make and the elements they contain
	//this creates a slice of 3 ints
	slice := make([]int, 3)
	fmt.Println(slice) //[0 0 0]

	//other builtins can be used in addition to slices
	//the builtin 'append' can be used to return a slice containg one or more new values
	slice = append(slice, 4)
	fmt.Println(slice) //[0 0 0 4]

	//slices can be copied from one to another
	//this creates a new slice equal to the length of the original slice
	//using the copy builtin we then copy the conents to the new slice
	copiedSlice := make([]int, len(slice))
	copy(copiedSlice, slice)
	fmt.Println(copiedSlice) //[0 0 0 4]

	//Slices support the slice operator to optain a range of vaules from within the slice
	newSlice := slice[3:]
	fmt.Println(newSlice) //[4]

	newSlice = slice[:3]
	fmt.Println(newSlice) //[0 0 0]

	//a slice can be declared and initialised in one line as well as arrays
	s := []string{"a", "b", "c"}
	fmt.Println(s) //[a b c]

}
