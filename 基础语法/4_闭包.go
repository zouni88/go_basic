package main

import "fmt"

func forward() func() int {

	return func() int {
		return 123
	}

}
func main() {
	//闭包调用，返回一个函数
	f := forward()
	fmt.Println(f())

}
