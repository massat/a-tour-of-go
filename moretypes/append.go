package main

import (
	"fmt"
)

func main() {

	var a []int
	printSlice("a", a)

	printSlice("a++", append(a, 5))
	printSlice("a++", append(a, 6, 7, 8, 9))
	printSlice("a++", append(a, 5, 6, 7, 8, 9, 1000))

	p := &a
	printSlice("p++", append(*p, 6, 7, 8, 9))
	printSlice("p++", append(*p, 5, 6, 7, 8, 9, 1000))

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
