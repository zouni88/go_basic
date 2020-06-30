package main

import "fmt"

func AnyFunc(v interface{}) {

}

func main() {
	var a interface{}
	var c *int
	//if x,y := c.(*int);y{
	//	a = *c
	//}
	a = c

	fmt.Println()
	fmt.Println(a)
}
