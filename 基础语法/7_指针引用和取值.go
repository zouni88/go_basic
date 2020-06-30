package main

import "fmt"

type persons struct {
	name string
}

func modify_name(name *persons) {
	name.name = "wang"
	per := *name
	fmt.Println("wodetian", per)
	fmt.Println(per.name)
}

func main() {
	person1 := persons{name: "cao"}
	//函数参数中接收者 带* 的参数，就是引用，不带* 就是直接取值，
	modify_name(&person1)

	//modify_name(person1)
	fmt.Println(person1.name)
	per := new(*persons)
	*per = &person1
	fmt.Println(per)
}
