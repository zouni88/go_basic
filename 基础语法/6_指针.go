package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	var a *int
	b := 123
	a = &b
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(&b)

	fmt.Println(*a)

	aa := 1
	bb := 2
	swap(&aa, &bb)
	fmt.Println(aa, bb)

}
