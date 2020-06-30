package main

import (
	"fmt"
	"reflect"
	"strings"
)

//是否包含字符串
func mconstains() {
	str := "12345"
	res := strings.Contains(str, "5")
	fmt.Println(res)

}

// replace 替换字符串
func mreplace() {
	str := "我的天哪我的天天天"
	//n 提换次数，-1 全部替换
	re := strings.Replace(str, "天", "地面", -1)
	fmt.Println(re)

}

//去掉字符串首尾的字符串
func mtrim() {
	str := " 我的天    那这是什么呀----"
	res := strings.Trim(str, " ")
	fmt.Println(res)
}

func mfeild() {
	//去除字符串中的空格，分割字符串变成切片，
	str := "我 了ge 咯咯 咯 咯咯 咯咯"
	res := strings.Fields(str)
	types := reflect.TypeOf(res)
	fmt.Println(res)
	//结果是切片类型
	fmt.Println(types)
}

func main() {
	mconstains()
	mreplace()
	mtrim()
	mfeild()

}
