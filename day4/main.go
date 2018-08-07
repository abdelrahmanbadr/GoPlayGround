package main

import (
	"fmt"
)

// we do this to declare behavior around a built-in type
type IP []byte

// type that represents the duration of time
// The type takes its representation from the built-in type int64
type Duration int64
type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}
type admin struct {
	person user // Embedded Type
	level  string
}

func struct_declaration() {
	lisa := user{"Lisa", "lisa@email.com", 123, true}
	// Pointers of type user
	lisa1 := &user{"Lisa", "lisa@email.com", 123, true}
	lisa1.notify()

	fred := admin{
		level:  "super",
		person: lisa,
	}
	fmt.Println(fred)
}

//Values of two different types can’t be assigned to each other, even if they’re compatible.
func duration_dec() {
	//dur declared and set to its zero value
	var dur Duration
	// will cause cannot error use int64(1000) (type int64) as type Duration in assignment
	dur = int64(1000)
}
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}
func main() {
	// Pointers of type user
	lisa1 := &user{"Lisa", "lisa@email.com", 123, true}
	fmt.Println(lisa1)
}
