package main

import (
	"fmt"
)

func main() {
	a := make([]int, 5)

	printSlice("a", a)

	b := make([]int, 0, 5)

	printSlice("b", b)

}

func printSlice(name string, slice []int) {
	fmt.Printf("Name: %s, len: %d, cap: %d %v \n", name, len(slice), cap(slice), slice)
}
