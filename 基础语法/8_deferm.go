package main

import "fmt"

func main() {
	defer fmt.Println("123")
	defer fmt.Println("444")
}
