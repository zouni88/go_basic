package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		for{
			fmt.Println("准备写入")
			ch <- 1
			fmt.Println("已经写入")
		}
	}()
	for v := range ch{
		fmt.Println(v)
	}

}
