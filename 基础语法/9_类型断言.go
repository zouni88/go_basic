package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i interface{}
	i = 123

	value,ok := i.(int)
	fmt.Println(value,ok)

	itype := reflect.TypeOf(i)
	fmt.Println(itype)
}
