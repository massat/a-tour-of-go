package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("There is %g problems", math.Nextafter(2, 3))
}
