package main

import (
	"fmt"
	"strconv"
)

//字符串转换

func main() {
	str := "ABCDEFGabcdefg"
	as := []byte(str)
	fmt.Println(as)

	// base 进制  2-36进制
	str = strconv.FormatUint(123, 2)
	fmt.Println(str)
	//float   fmt: 打印格式 'f' 小数方式，prec: 小数位
	b := 123.1234
	str = strconv.FormatFloat(b, 'f', 2, 32)
	fmt.Println(str)

	// 直接转换整型为字符串
	str = strconv.Itoa(1239)
	fmt.Println(str)

	res, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(res)
	} else {
		fmt.Println(res, err)
	}

}
