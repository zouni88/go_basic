package main

import (
	"container/list"
	"fmt"
)

func main() {
	lis := list.New()

	for i := range [10]int{} {
		lis.PushBack(i)
	}
	for i := lis.Front(); i != nil; i = i.Next() {
		fmt.Println(i)
	}

	fmt.Println(lis.Front())
	fmt.Println(lis.Front())
	i := lis.Front().Next().Next().Next().Next()
	fmt.Println(i)
}
