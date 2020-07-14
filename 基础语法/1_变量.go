package main

import "fmt"

var AAA struct {
	Id int `json:"id" xml:"id"`
}

func main() {
	var a = 1 //var显式声明
	b := 1    //智能推断
	fmt.Println(a, b)
	var c interface{}
	c = 123
	fmt.Println(c)

	//变量引用
	var aa = 123
	func(aa int) {
		aa = 456
	}(aa)
	fmt.Println(aa)
}
