package main

import "fmt"

func main() {
	c := make(chan bool)
	go func() {
		c <- true
		fmt.Println("放进去了")
	}()
	select {
	case <- c:
		fmt.Println("结束")
	}
}
