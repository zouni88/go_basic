package main

import (
	"fmt"
	"sync"
)

var count = 10
var lock sync.Mutex

func Consume(wg *sync.WaitGroup)  {

	lock.Lock()
	defer wg.Done()
	defer lock.Unlock()
	if count == 0{
		fmt.Println("没有库存了")
		return
	}
	count --

	fmt.Println(count)
}

func main() {
	wg := sync.WaitGroup{}
	for range [100]int{}{
		wg.Add(1)
		go Consume(&wg)
	}
	wg.Wait()
	fmt.Println(count)
}
