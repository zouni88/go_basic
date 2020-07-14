package main

import "fmt"

func divide(a int,b int,isPanic bool) int  {
	//最终执行 recover ,类似于 java try catch ，，通过panic抛出异常
	if isPanic{
		defer func() {
			if err := recover(); err != nil{
				fmt.Println(err)
			}
		}()
	}
	if b==0{
		panic("除数不能等于0")
	}
	c := a/b
	return c
}

func main() {
	divide(3,0,true)
	fmt.Println("我的天，一切运转正常")
}


