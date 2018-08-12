package main

import (
	"fmt"

	. "./singleton"
)

func main() {
	instance := GetInstance()
	fmt.Println(instance.Increment())
	fmt.Println(instance.Increment())
	new := GetInstance()
	fmt.Println(new.Increment())
}
