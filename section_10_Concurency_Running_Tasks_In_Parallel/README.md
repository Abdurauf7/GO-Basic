package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

// chanelName <- value (here we are sending value to chanel)
func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

// go -> keyword make function run paralel

//done := make(chan bool)  --> creating channels(like comunication device)

// <-done (receiving data to display it)

func main() {
	dones := make([]chan bool, 4)

	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!", done)

    <-done
	<-done
	<-done
	<-done
}

==================== Multiple wat of doing it================

func main() {
	dones := make([]chan bool, 4)

	dones[0] = make(chan bool)
	go greet("Nice to meet you!", dones[0])

	dones[1] = make(chan bool)
	go greet("How are you?", dones[1])

	dones[2] = make(chan bool)
	go slowGreet("How ... are ... you ...?", dones[2])

	dones[3] = make(chan bool)
	go greet("I hope you're liking the course!", dones[3])

	for _, done := range dones {
		<-done
	}
}

==========================Third way of doing it===================================
package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

// chanelName <- value (here we are sending value to chanel)
// close(chanelName) -- closing chanel
func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan)
}

// go -> keyword make function run paralel

//done := make(chan bool)  --> creating channels(like comunication device)

// <-done (receiving data to display it)

func main() {
	// dones := make([]chan bool, 4)

	done := make(chan bool)
	// dones[0] = make(chan bool)
	go greet("Nice to meet you!", done)
	// dones[1] = make(chan bool)
	go greet("How are you?", done)
	// dones[2] = make(chan bool)
	go slowGreet("How ... are ... you ...?", done)
	// dones[3] = make(chan bool)
	go greet("I hope you're liking the course!", done)

	

	for range done {

	}
}
