package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{Y: 10}
	v3 = Vertex{}
	p  = &Vertex{10, 30}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
