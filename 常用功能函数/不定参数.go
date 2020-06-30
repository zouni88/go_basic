package main

import "fmt"

func args(args ...int) {
	for i, v := range args {
		fmt.Println(i, v)
	}
}

func sum(args ...int) {
	sum := 0
	for _, v := range args {
		sum += v
	}
	fmt.Println(sum)
}

func main() {
	args(1, 2, 3, 4)
	sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
}
