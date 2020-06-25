package main

import "fmt"

func divi(a,b int )(res int,err error)  {
	if b==0{
		fmt.Println(err)
	} else {
		res = a/b
	}
	return
}

func main() {
	res,err := divi(1,0)
	fmt.Println(res,err)
}
