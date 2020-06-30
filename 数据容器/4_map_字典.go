package main

import "fmt"

func main() {
	ma := map[int]string{1: "value", 2: "value1"}
	for i, v := range ma {
		fmt.Println(i, v)
	}
	fmt.Println(ma)
	fmt.Println(ma[1])
}
