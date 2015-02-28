package main

import (
	"fmt"
)

func main() {
	pow := make([]int, 10)

	for i, v := range pow {
		fmt.Println(i, v)

		pow[i] = 1 << uint(i)
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

}
