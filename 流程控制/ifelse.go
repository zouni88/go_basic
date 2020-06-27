package main

import (
	"fmt"
	"os"
)

func main() {
	a := true
	if a {
		fmt.Println(a)
	}
	switch a {
	case true:
		fmt.Println("是的")
	case false:
		fmt.Println("波士顿")
	}

	b := "str"
	switch b {
	case "str","sub":
		fmt.Println("我的天")
	case "wo":
	}
	

}
