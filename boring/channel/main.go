// A channel allows for communication and synchronization between goroutines.
//
// Level: beginner
// Topics: goroutines, channels
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	go say("blah")
	for i := 0; i < 1000; i++ {
		fmt.Println("Waiting")

		time.Sleep(time.Millisecond * 500)

	}
}

func say(msg string) {
	for i := 0; ; i++ {
		fmt.Printf("%s, %d \n", msg, i)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))

	}
}
