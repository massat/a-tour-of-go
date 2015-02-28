package main

import (
	"fmt"
)

func main() {
	s := []int{2, 4, 6, 8, 10}
	fmt.Println("s ==", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("p[%d] == %d\n", i, s[i])
	}
}
