package main

import (
	"sync"
	"time"
)

var phoneCount = 10
var capac = make(chan int,10)
var wg = sync.WaitGroup{}
func main() {
	for k,_ := range [10]int{}{
		wg.Add(1)
		go request(k)
	}
	go dealCapc()
	wg.Wait()
}

func dealCapc(){
	println("库存剩余：[%d]",phoneCount)
	for  {
		uid:= <- capac
		phoneCount--
		println("用户id:[%d] 抢到一台，手机编号：[%d]",uid,phoneCount)
	}
}

func server(uid int)  {
	capac <- uid
}

func request(uid int)  {
	time.Sleep(time.Second)
	server(uid)
}
