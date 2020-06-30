package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i interface{}
	i = 123
	fmt.Println(i)
	var a []interface{}
	a = append(a, "asdf", func() { fmt.Println("1234func") })
	for i, v := range a {
		types := reflect.TypeOf(v)
		fmt.Println(types)
		_, a := v.(func())
		fmt.Println(i, v, a)

	}

}
