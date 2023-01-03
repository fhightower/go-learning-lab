package main

import (
	"fmt"
)

func f1(a int) int {
	defer f2()
	return f3()
}

func f2() {
	fmt.Println("f2")
}

func f3() int {
	fmt.Println("f3")
	return 1
}

func main() {
	f1(1)
}

// f3
// f2
