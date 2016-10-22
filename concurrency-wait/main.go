package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printHello(i int) {
	fmt.Println("Hello", i)
	// finish the wait group
	wg.Done()
}

func main() {

	for i := 0; i < 5; i++ {
		// create a new wait group
		wg.Add(1)
		go printHello(i)

	}
	// wait until all wg are done
	wg.Wait()

}
