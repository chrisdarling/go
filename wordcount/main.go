package main

import (
	"fmt"
	"strings"
)

func wordCount(str string) int {
	// strings.Fields(str) splits the string around each instance of whitespace
	return len(strings.Fields(str))

}

func main() {
	str := "Put   some       words here"
	fmt.Println(wordCount(str))

}
