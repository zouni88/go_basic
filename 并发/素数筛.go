package main

import (
	"fmt"
)

func GenerNatural() (chan int) {
	ch := make(chan int)
	go func() {
		for i:= 2; ;i++{
			ch <- i
		}
	}()
	return ch
}

func PrimeFilter(in <-chan int, prime int) (chan int) {

	out := make(chan int)
	go func() {
		for {
			if i:= <-in; i%prime != 0{
				out <-i
			}
		}
	}()
	return out
}

func main() {
	ch := GenerNatural()
	fmt.Println(&ch)

	for i:= 0;i<10;i++{
		prime := <-ch
		fmt.Println(prime)
		ch = PrimeFilter(ch,prime)
	}
}
