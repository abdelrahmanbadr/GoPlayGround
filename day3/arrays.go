package main

import (
	"fmt"
)

func array_declaration() {
	var arr1 [2]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [5]int{0: 5, 4: 3}
	arr4 := [...]int{9}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)

}

func array_of_pointers() {
	//An array of pointers that point to integers
	array := [3]*int{new(int), new(int), new(int)}
	*array[0] = 10
	*array[1] = 20
	fmt.Println(*array[1])
}

//Assigning one array to another of the same type
func array_assiging() {
	array1 := [5]string{"A", "B", "C", "D", "E"}
	array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	array1 = array2
	fmt.Println(array1)

}

func multi_array() {
	array := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	fmt.Println(array)
	fmt.Println(array)
}
func foo(array *[2]int) {
	fmt.Println(array[0])
}
