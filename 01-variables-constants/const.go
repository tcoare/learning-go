package main

import (
	"fmt"
	"math"
)

const x string = "constant" //"const" is used to declare a constant value

func main() {

	//declare variables using var
	var a = "String"
	fmt.Println(a) //String

	//can declare more than one variable at once
	var b, c int = 1, 2
	fmt.Println(b, c) //1 2

	//Go can infer the type of an initialised variable
	var d = 5
	fmt.Println(d) //5

	//Go will zero value a variable if declared with no value
	var e int
	fmt.Println(e) //0

	//Short hand assignment where var can be omitted and type is still infered
	//can only be used within a function
	f := "Hello World"
	fmt.Println(f) //Hello World

	fmt.Println(x) //constant

	//"const" can be used where a "var" can be used
	const num = 500000

	const y = 3e20 / num
	fmt.Println(y)

	//a numeric constant has no type until its given one, this can be done by an explicit conversion
	fmt.Println(int64(y))

	//or using it in a context that would require one. In this case 'math.Sin' expects a float64
	fmt.Println(math.Sin(num))

}
