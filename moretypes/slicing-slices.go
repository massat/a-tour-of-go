package main

import (
	"fmt"
)

func main() {
	p := []int{2, 3, 5, 8, 8, 10}
	fmt.Println("p ==", p)
	fmt.Println("p[1:4] ==", p[1:4])

	fmt.Println("p[:3] ==", p[:3])
	fmt.Println("p[3:] ==", p[3:])
}
