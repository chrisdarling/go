package main

import "fmt"

// recieve channel store in message and print
// sends ping through the channel
func pinger(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		c <- "ping"
	}

}

// recieves channel store in message and print
// sends pong through the second channel
func ponger(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		c <- "pong"
	}

}

func main() {

	// create a chanels
	c := make(chan string)

	// order of calls do not matter in this case
	// go routines will wait until channel is sent because they recive first
	go ponger(c)
	go pinger(c)

	// start by send the channel
	c <- "Start"

	// enter into command line to stop ping pong
	var input string
	fmt.Scanln(&input)
}
