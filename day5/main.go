package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func main() {
	Playerchannel()
	// Allocate 1 logical processor for the scheduler to use.
	// runtime.GOMAXPROCS(5)
	//A WaitGroup is a counting semaphore that can be used to maintain a record of running goroutines

	wg.Add(2)
	fmt.Println("Start Goroutines")
	// Declare an anonymous function and create a goroutine.
	go func() {
		// Schedule the call to Done to tell main we are done.
		defer wg.Done()
		// Display the alphabet three times
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()
	// Declare an anonymous function and create a goroutine.
	go func() {
		// Schedule the call to Done to tell main we are done.
		defer wg.Done()
		// Display the alphabet three times
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()
	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")
}

func raceCondition() {
	// Add a count of two, one for each goroutine.
	wg.Add(2)
	// Create two goroutines.
	go incCounter(1)
	go incCounter(2)
	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

// incCounter increments the package level counter variable.
func incCounter(id int) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()
	for count := 0; count < 2; count++ {
		// Capture the value of Counter.
		value := counter
		// Yield the thread and be placed back in queue.
		runtime.Gosched()
		// Increment our local value of Counter.
		value++
		// Store the value back into Counter.
		counter = value
	}

}
func mutx() {
	// Add a count of two, one for each goroutine.
	wg.Add(2)
	// Create two goroutines.
	go incCounter_mutex(1)
	go incCounter_mutex(2)
	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Printf("Final Counter: %d\\n", counter)
}

// incCounter increments the package level Counter variable
// using the Mutex to synchronize and provide safe access.
func incCounter_mutex(id int) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()
	for count := 0; count < 2; count++ {
		// Only allow one goroutine through this
		// critical section at a time.
		mutex.Lock()
		{
			// Capture the value of counter.
			value := counter
			// Yield the thread and be placed back in queue.
			runtime.Gosched()
			// Increment our local value of counter.
			value++
			// Store the value back into counter.
			counter = value
		}
		mutex.Unlock()
		// Release the lock and allow any
		// waiting goroutine through.
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
