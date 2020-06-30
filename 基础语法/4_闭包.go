package main

import "fmt"

func forward(value int) func() int {
	return func() int {
		value ++
		return value
	}

}
func main() {
	//闭包调用，返回一个函数

	res := forward(1)
	fmt.Println(res())
	fmt.Println(res())
	fmt.Printf("%p\n", res)
	res1 := forward(10)
	fmt.Println(res1())
	fmt.Println(res1())

	ress, err := fmt.Printf("%p,%p\n", res,res1)
	fmt.Println(ress,err)


	//e := 123
	//c := func() {
	//	e = 666
	//}
	//c()
	//fmt.Println(e)
	a := 123
	res111 := fmt.Sprintf("我的%v",a)
	fmt.Println(res111)


}
