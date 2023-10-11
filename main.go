package main

import (
	"fmt"
	"math"
)

func main() {
	number := 25.0
	squareRoot := math.Sqrt(number)

	fmt.Printf("The square root of %.2f is %.2f\n", number, squareRoot)
}
