package main

import (
	"fmt"
)

func slice_declaration() {
	// Slice contains a length of 3 and has a capacity of 5 elements.
	slice1 := make([]int, 3, 5)
	// Slice contains a length and capacity of 5 elements.
	slice2 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}

	// Initialize the 100th element with an empty string.
	slice3 := []string{99: ""}
	// Create a nil slice of integers.
	// Memory isn't allocated yet
	var nil_slice []int
	// Use a slice literal to create an empty slice of integers.
	// Memory allocated
	empty_slice := []int{}

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println("nil slice", nil_slice)
	fmt.Println("empty slice", empty_slice)
}

func slice_subset() []string {
	s1 := []string{"Red", "Blue", "Green", "Yellow", "Pink", "Orange"}
	// slice points to index number 2 in the original array
	// slice start with index number 2 and ends with index number 3
	// s2 length = 2 but capacity = 4 because it points to the original array
	// s2 now points to original array starting from index 2 {"Green", "Yellow", "Pink", "Orange"}
	s2 := s1[2:4]
	s2[1] = "Black"
	fmt.Println(s2)
	// Notice that s1 changed when s2 changed
	// Because they are both points to the same array
	fmt.Println(s1)

	new_slice := make([]string, 5)
	copy(new_slice, s1)
	s1[1] = "White"
	// Notice that new_slice haven't changed
	// Because it's point to different array
	fmt.Println(new_slice)
	fmt.Println(s1)
	// using ... because its variadic function (... should be used with arrays and slices)
	s1 = append(new_slice, s1...)
	return s2
}
func solve_memoryLeak_sub_slice() []int {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s2 := make([]int, 3)
	// when you copy to new slice, you are allocating to new array
	// when the function exit the grapage collector will now that s1 is not used any more
	copy(s2, s1[1:4])
	return s2
}
func remove() []int {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	index := 2 //delete index number 2
	s1 = append(s1[:index], s1[index+1:]...)
	return s1
}

// you can't access index out of range even it was in slice capacity
func slice_capacity_error() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// s2 capacity is 7 , s2 length = 2
	s2 := s1[1:3]
	//will show [2 3 4 5 6 7 8]
	fmt.Println(s2[:7])
	//this will cause runtime error
	fmt.Println(s2[5])
	fmt.Println(s2[5:])
}
func three_index_slice() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// Slice the third element and restrict the capacity.
	// Contains a length of 1 element and capacity of 2 elements.
	s2 := s1[2:3:4]
	fmt.Println("capacity is :", cap(s2), "\nit points to: ", s2[:cap(s2)])

}
func multidimensional_slice() {
	slice := [][]int{{10}, {100, 200}}
	fmt.Println(slice)
}
