package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4}
	//var c *[5]int = &a
	c := &a
	c[0] = 222
	fmt.Println(*c)
	fmt.Printf("%p\n", &a)
	fmt.Println(&(a[1]))
	fmt.Println(&c[0])
	fmt.Println(&c[1])
	fmt.Println(*c)

	//var b *[3]int
	//*b[0] = &a[0]
	//fmt.Printf("%p",&b)

}
