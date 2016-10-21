package main

import (
	"fmt"
	"strings"
)

func main() {
	// convert to upper case
	fmt.Println(strings.ToTitle("loud noises!!!"))
	fmt.Println(strings.ToLower("LOW NOISES ..."))

	// replace name with chris once
	fmt.Println(strings.Replace("Hello name?", "name", "Chris", 1))

	// remove spaces
	fmt.Println(strings.TrimSpace("  remove spaces    "))

}
