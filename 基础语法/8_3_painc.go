package main

import "fmt"

func main() {
	a := 1
	b := 0

	if b == 0{
		//panic 配合recover使用
		panic("除数不能等于0")
	}
	c := b/a
	fmt.Println(c)
}
