package main

import (
	"errors"
	"fmt"
)

type ErrorCustom interface {
	Error()
}

func Add(a int,b int)  (int,error){
	var res int
	var err error
	if a <0{
		err = errors.New("不能小于0")
	} else {
		res = a+b
	}
	return res,err
}

func Division(a int,b int) (int,error) {
	var err error
	if b <= 0{
		err = errors.New("除数不能小于0")
		return 0,err
	}
	return a/b,err

}
func main()  {
	res,err := Add(1,2)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(res)

	res1,err1 := Division(1,0)
	fmt.Println(res1,err1)
}