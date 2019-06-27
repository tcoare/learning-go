package main

import "fmt"

func main() {

	//create a map using the 'make' builtin
	userDB := make(map[string]int)

	//key and values are associated using 'name[key] = value'
	userDB["Tyler"] = 293856
	userDB["John"] = 572892
	userDB["Test"] = 0

	fmt.Println(userDB) //map[John:572892 Tyler:293856]

	//can obtain a value of a key using 'name[key]'
	user1ID := userDB["Tyler"]
	fmt.Println(user1ID) //293856

	//the delete builtin can be used to remove pairs from a map
	delete(userDB, "John") //delete(map, key)
	fmt.Println(userDB)    //map[Tyler:293856]

	//the second return value when getting a value of a key is its presence in the map
	//we dont need the value of the key so we use a blank identifier '_'
	//this is used to differentiate between 0 value keys and keys which aren't present
	_, present := userDB["John"]
	fmt.Println(present) //false

	_, present = userDB["Test"]
	fmt.Println(present) //true

	//maps can be declared and initialised on the same line
	//omit the 'make' builtin and populate using {key1: val1, key2: val2, ...}
	countryPopulation := map[string]int{"England": 55000000, "United States": 327000000}
	fmt.Println(countryPopulation) //map[England:55000000 United States:327000000]
}
