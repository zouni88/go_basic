package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		for i, v := range [10]int{} {
			fmt.Println(i, v)
		}
	}()
	//没有下面的这句，程序什么都不会输出， go 主线程不会等待协程，
	time.Sleep(time.Second)
	fmt.Println("睡眠充足...")
}
