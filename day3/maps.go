package main

import (
	"fmt"
)

func check_map(data map[string]int) {
	fmt.Println(data)
	// fmt.Println(data["ay_7aga"]) // will resut 0 because it doesn't exit
	if val, Ok := data["ay_7aga"]; Ok {
		fmt.Println(val)
	} else {
		//will result false if doesn't exist
		fmt.Println(Ok)
	}

}

//a map is an unordered collection of key/value pairs
func map_declaration() {
	// dict := make(map[string]int)

	// Create a map using a slice of strings as the value.
	// dict := map[int][]string{}
	m := map[string]int{
		"Ahmed": 23,
		"Sara":  20,
	}
	// Remove the key/value pair for the key "Coral".
	delete(m, "Ahmed")
	check_map(m)
}
