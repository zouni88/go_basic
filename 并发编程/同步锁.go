package main

import (
	"fmt"
	"sync"
	"time"
)

var count = 10
var lock sync.Mutex

func Consume()  {
	lock.Lock()
	if count == 0{
		fmt.Println("没有库存了")
		return
	}
	count --
	lock.Unlock()
	fmt.Println(count)
}

func main() {

	for range [100]int{}{
		go Consume()
	}
	time.Sleep(time.Second)
	fmt.Println(count)
}
