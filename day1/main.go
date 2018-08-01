package main

import (
	"errors"
	"fmt"
)

type Student struct {
	Age  int
	Name string
}

func NewStudent() Student {
	return Student{Age: 18, Name: "ahmed"}
}

func main() {
	fmt.Println(NewStudent())
	x := 0
	y := -1
	var sum int = x + y
	result, err := sumSlice(sum)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	mapData()

}

func sumSlice(len int) (int, error) {
	if len < 0 {
		return 0, errors.New("no negative number allowed")
	}
	var slice = make([]int, len)
	sum := 0
	for i := 1; i <= len; i++ {
		slice[i-1] = i
	}
	for i := 0; i < len; i++ {
		sum += slice[i]
	}
	return sum, nil
}

func mapData() {
	studentAges := make(map[string]int)
	studentAges["st1"] = 22
	studentAges["st2"] = 18
	studentAges["st3"] = 20
	for _, val := range studentAges {

		fmt.Println(val)
	}
	delete(studentAges, "st2")
}
