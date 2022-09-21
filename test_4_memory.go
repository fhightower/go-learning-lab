package main

import (
	"fmt"
)

func main() {
	// long way:
	// var i int = 5
	// short way using type inference:
	i := 5
	fmt.Println(i)
	fmt.Println(&i)

	n := 1
	inc_var(n)
	fmt.Println(n)
	inc_mem_loc(&n)
	fmt.Println(n)
}

func inc_var(n int) {
	n++
}

func inc_mem_loc(n *int) {
	*n++
}

