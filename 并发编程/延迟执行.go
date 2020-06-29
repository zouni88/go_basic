package main

import (
	"fmt"
	"sync"
	"time"
)

var name string = ""
func C()  {
	time.Sleep(time.Second*4)
	fmt.Println("程序执行结果呢。。。",name)
}


var Once sync.Once
func main() {
	go C()
	fmt.Println("结束了")
	name = "Once"
	Once.Do(C)
	// Once 执行函数指挥执行一次，并且主线程会等待执行结果，


}
