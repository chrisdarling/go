package main

import (
	"fmt"
	"time"
)

func main() {
	timeAsString := "02/03/2016"
	//format is a date of how you want it too look
	// 2006 where date will be
	// 01 where month will be
	// 02 where day will be
	newTime, _ := time.Parse("02/01/2006", timeAsString)
	// Parse will default to UTC you can format it by usinf
	// time.Format function
	fmt.Println(newTime.Format("02/01/2006"))
}
