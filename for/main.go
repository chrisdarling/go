package main

import (
	"fmt"
)

func main() {
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	for x := 0; x < 5; x++ {
		if x%2 == 0 {
			fmt.Println(x, ": even")
		} else {
			fmt.Println(x, ": odd")
		}
	}

	// No break statement in Go!!!!

	for y := 0; y < 5; y++ {
		switch y {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		default:
			fmt.Println("Unknown Number")
		}
	}

}
