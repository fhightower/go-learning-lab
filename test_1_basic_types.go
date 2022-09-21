package test_1

import (
	"fmt"
)

func main() {
	// var s string = "Foo bar"
	s := "Foo bar"
	fmt.Println(s)

	// var a [5]int
	// a[2] = 7
	// fmt.Println(a)

	a := [5]int{0, 1, 2, 3, 4}
	fmt.Println(a)

	slice := []int{0, 1, 2, 3, 4}
	slice = append(slice, 27)
	fmt.Println(slice)

	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	fmt.Println(m)
	delete(m, "a")
	fmt.Println(m)
	fmt.Println(m["c"])

	// basic for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// iterate over array
	for index, value := range a {
		fmt.Println(index, value)
	}

	// iterate over map
	for key, value := range m {
		fmt.Println(key, value)
	}
}
