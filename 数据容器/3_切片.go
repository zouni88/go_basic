package main

import "fmt"

func main() {
	//切片不用指定大小，相当于动态的数组
	lis := []int{1,2,3,3}
	for i,v := range lis{
		fmt.Println(i,v)
	}
	fmt.Println(lis)
}
