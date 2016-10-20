package main

import (
	"fmt"
)

func count(x []float64) float64 {
	total := 0.0
	for _, v := range x {
		total += v
	}

	return total
}

func main() {
	mySlice := []float64{32, 54, 75, 78, 756}
	fmt.Println(count(mySlice))
}
