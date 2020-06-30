package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := [3]int{4, 5, 6}
	// 指针数组内的每个元素都是地址，指向元素地址的变量
	var aa = [2]*[3]int{&a, &b}
	fmt.Println(aa)
	fmt.Printf("%p,%p", &a, aa[0])

	//var a := *[3]int{}

	e := [][][]int{{{1}}}
	fmt.Println(e)

}
