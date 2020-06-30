package main

import "fmt"

func main() {
	var a int
	fmt.Printf("%p", &a)
	_, _ = fmt.Scanf("%d", &a)

	fmt.Printf("%v", a)

}
