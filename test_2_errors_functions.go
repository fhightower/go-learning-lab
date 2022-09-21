package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sqrt(16))
	fmt.Println(sqrt(-1))
}

func sum(a int, b int) int {
	return a + b
}

func sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("Undefined")
	}

	return math.Sqrt(a), nil
}

