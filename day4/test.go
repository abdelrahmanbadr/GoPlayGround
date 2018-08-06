package main

import "fmt"

type notifier interface {
	notify()
}

// func newUser() notifier {
// 	return users{}
// }

type users struct {
	name  string
	email string
}

func (u *users) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

func main() {
	// Create a value of type User and send a notification.
	u := users{"Bill", "bill@email.com"}
	sendNotification(&u)
}
func sendNotification(n notifier) {
	n.notify()
}
