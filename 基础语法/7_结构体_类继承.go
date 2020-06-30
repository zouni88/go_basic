package main

import "fmt"

type People struct {
	id  int
	age int
}

type Student struct {
	People
	name string
}

func main() {
	var wang Student
	wang.id = 1
	wang.age = 18
	wang.name = "zhang"
	fmt.Println(wang)
}
