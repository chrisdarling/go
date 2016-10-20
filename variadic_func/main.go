// you can have a function that takes zero or more functions
// thesse are called variadic functions
package main

import (
	"fmt"
)

// ... specifies it can take more than on argument
// then you define the type
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}

	return total
}

func main() {
	fmt.Println(add(1, 2, 3, 4, 5))
}
