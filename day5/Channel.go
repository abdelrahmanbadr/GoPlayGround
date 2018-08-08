package main

import (
	"fmt"
	"math/rand"
)

// main is the entry point for all Go programs.
func Playerchannel() {
	// Create an unbuffered channel.
	court := make(chan int)
	// Add a count of two, one for each goroutine.
	wg.Add(2)
	// Launch two players.
	go player("Nadal", court)
	go player("Djokovic", court)
	// Start the set.
	court <- 1
	// Wait for the game to finish.
	wg.Wait()
}
func player(name string, court chan int) {
	defer wg.Done()
	for {
		ball, Ok := <-court
		if !Ok {
			fmt.Println("Player Name ", name)
			return
		}
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			close(court)
			return
		}
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		court <- ball
	}
}
