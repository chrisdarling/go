// generate html to stdout and save it to a file
// "Enter command go run main.go <number to convert> > <filename>.html"
// Your html will be generated at the path you specify

package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	miTokm = 1.60934
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Enter a number as an argument!")
		os.Exit(1)
	}
	// if there number entered is not a float number will be set to 0
	number, _ := strconv.ParseFloat(os.Args[1], 64)

	fmt.Println("<!Doctype html>")
	fmt.Println("<html>")
	fmt.Println("<head></head>")
	fmt.Println("<body>")

	fmt.Printf("Miles: %3.2f\n <br />", number)

	fmt.Printf("Kilometers: %3.2f\n", number*miTokm)
	fmt.Println("</body>")
	fmt.Println("</html>")
}
