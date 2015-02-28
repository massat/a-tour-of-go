package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	var today int = time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today is Saturday!")
	case today + 1:
		fmt.Println("Tomorrow")
	}
}
