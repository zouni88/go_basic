package main

import "fmt"

var AAA struct {
	Id int
}

func main() {
	var a = 1 //var显式声明
	b := 1    //智能推断
	fmt.Println(a, b)
	var c interface{}
	c = 123
	fmt.Println(c)
}
