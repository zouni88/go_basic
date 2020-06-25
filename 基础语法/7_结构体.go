package main

import "fmt"

type student struct {
	id int
	name string
}
func main() {
	stu := student{1,"cao"}
	fmt.Println(stu.id)
	//结构体指向第一个变量地址
	fmt.Printf("%p\n",&stu)
	fmt.Printf("%p\n",&stu.id)
	fmt.Printf("%p\n",&stu.name)

	a := make([]student,3)
	a[0] = student{1,"cao"}
	fmt.Println(a)

	stu1 := stu
	stu.name = "我的天"
	fmt.Println(stu.name,stu1.name)
}
