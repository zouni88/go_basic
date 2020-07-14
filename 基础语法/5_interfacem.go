package main

import (
	"errors"
	"fmt"
)

type Minter interface {
	sum(a, b int) int
}

func (mm Phone) sum(a, b int) int {
	fmt.Println(a + b)
	return a + b
}

func main() {
	addd := new(Phone)
	a := addd.sum(1, 2)
	fmt.Println(a)
	//aa:= 1
	//bb:= 2

	var err = errors.New("错误信息？？？")
	fmt.Println(err)

}

//func max(a,b int) (int,error)  {
//	if a<b{
//		return b,errors.New("错误信息")
//	}
//	//return 1
//}

type error interface {
	Error() string
}
