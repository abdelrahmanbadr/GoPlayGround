package main

import (
	"fmt"
)

// This function forces Car to implement all methods of Vehicle Interface
func NewCar() Vehicle {
	// return  Pointer of type car
	return &Car{}
}

type Vehicle interface {
	start_engine()
	drive()
}

type Car struct{}
type Truck struct{}

// start_engine implements a method with a pointer receiver.
func (c *Car) start_engine() {}
func (c *Car) drive()        {}

func (t *Truck) start_engine() {}
func (t *Truck) drive()        {}

//Polymorphism is the ability to write code that can take on different behavior through the implementation of types
func main() {
	//polymorphism is achived here
	var v Vehicle = &Car{}
	// start_engine method with pointer receiver
	// v is a pointer of type car
	v.start_engine()
	v.drive()
	/////
	accelerate_engine(v)
}
func accelerate_engine(v Vehicle) {
	v.start_engine()
	fmt.Println("accelerating")
	// do whatever to accelerate
}
