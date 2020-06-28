package main

import (
	"fmt"
	"reflect"
)

const(
	A = iota +1
	B
	C
)

func AB() int {
	return 123
}
func main() {
	fmt.Println(A,B,C)

	res := reflect.ValueOf(AB())
	fmt.Println(res)
}

