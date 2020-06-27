package main

import "fmt"

func main() {
	//最后执行
	defer fmt.Println("123")
	//倒数第二
	defer fmt.Println("444")
	//先执行
	for i:= range [50]int{}{
		fmt.Println(i)
	}
}
