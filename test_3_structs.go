package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	p := person{name: "Alice", age: 21}
	fmt.Println(p)
	fmt.Println(p.age)
}
