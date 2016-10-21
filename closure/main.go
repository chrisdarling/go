package main

import "fmt"

// you can create function inside of function in go
// This is a function that has no arguments and returns
// function that return an int

func incrementor(x int) func() int {
	return func() int {
		x++
		return x
	}
}

func main() {
	// setting numbers to the function incrementor
	// brings the scope of var x inside of incremetor to
	// the scope of function main
	// the result is x continues to be incremented
	numbers := incrementor(2)
	fmt.Println(numbers())
	fmt.Println(numbers())
	fmt.Println(numbers())

}
