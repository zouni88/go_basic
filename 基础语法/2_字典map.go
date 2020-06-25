package main

import (
	"fmt"
)

func main() {
	//var dic map[int]string
	dic := make(map[int]string)
	dic[1] = "cao"
	dic[2] = "wang"
	dic[3] = "zhang"
	for k,v := range dic {
		fmt.Println(k,v)
	}
	fmt.Println(dic)

	//删除map元素
	delete(dic,2)
	fmt.Println(dic)
	delete(dic,5)
	fmt.Println(dic)
	//diccopy = copy(dic,diccopy)

}