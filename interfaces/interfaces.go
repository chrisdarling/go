package main

import "fmt"

// declare a type square
type square struct {
	side float64
}

// declare a ype rect
type rect struct {
	len float64
	wid float64
}

// give square a function area
func (z square) area() float64 {
	return z.side * z.side
}

// give rect a function area
func (r rect) area() float64 {
	return r.len * r.wid
}

// declare an interface of type shape
// Any type that has a function area() that returns a float64
// is implicintly implementing the shape interface

// interfaces define behavior and allow polymorphism
type shape interface {
	area() float64
}

// declare a function that takes a shape as a parameter
func info(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
}

func main() {
	x := square{64}
	y := rect{20, 40}
	info(x)
	info(y)

}
