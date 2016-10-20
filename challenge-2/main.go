package main

import "fmt"

//Write a program that finds the smallest number in this list:

func smallest(x []int) {
	if x == nil {
		fmt.Println("Empty array")
		return
	}

	sm := x[0]
	for _, v := range x {
		if v < sm {
			sm = v
		}
	}

	fmt.Println(sm)
}

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	smallest(x)
}
