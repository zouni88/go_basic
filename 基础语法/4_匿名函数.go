package main

import "fmt"

func main() {
	add := func() (min int) {
		min = 123
		return
	}
	res := add()
	fmt.Println(res)

	//匿名函数返回值
	max, min := func(a, b int) (min, max int) {
		min = 0
		max = 255
		return
	}(0, 266)
	fmt.Println(max, min)

}
