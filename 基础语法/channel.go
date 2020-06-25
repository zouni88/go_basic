package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for{
			ch <- 1
			time.Sleep(time.Second)
		}
	}()
	for v := range ch{
		fmt.Println(v)
	}


}
