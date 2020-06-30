package main

import "fmt"

func main() {
	//数组是固定大小的，声明必须指定容量
	var lis = [5]int{1, 2, 3}
	for i, v := range lis {
		fmt.Println(i, v)
	}
	//输出结果 [1 2 3 0 0] ,未初始化的值 默认为 0
	fmt.Println(lis)

	b := byte('1')
	fmt.Println(b)

}
