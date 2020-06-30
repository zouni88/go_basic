package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%p\n", &arr)

	slice := arr[1:3]
	slice[0] = 10
	fmt.Println(slice)
	fmt.Println(arr)
	fmt.Printf("%T", slice)
	fmt.Printf("%p\n", slice)

}
