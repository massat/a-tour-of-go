package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2700

	p := &i
	fmt.Println(p)
	fmt.Println(*p)

	p = &j
	*p = *p / 37
	fmt.Println(*p)
}
