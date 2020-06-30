package main

import "fmt"

func main() {
	a := 123
	fmt.Printf("%v\n", a)

	var b *int
	b = &a
	fmt.Printf("%p\n", b)
	fmt.Printf("%v\n", &b)
	fmt.Println(*b)

	fmt.Printf("%b", 129)
}
