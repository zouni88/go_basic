package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func basicoper() {
	f, r := os.Create("abc.txt")
	ftype := reflect.TypeOf(f)
	rtype := reflect.TypeOf(r)
	fmt.Println(ftype)
	fmt.Println(rtype)
	f.WriteString("我的天")
	n, e := f.Write([]byte("我的天"))
	fmt.Println(n, e)

	//fmt.Println(s,e)
	//s,e := os.Getwd()
	defer f.Close()
}

func writefile() {
	file, e := os.Create("aa.txt")
	if e != nil {
		fmt.Println("错误了")
		return
	}
	file.WriteString("我的天哪")
	file.WriteString("这是谁呀")

	r, e := file.Seek(0, io.SeekEnd)
	fmt.Println(r, e)

	file.WriteAt([]byte("哈哈哈"), 0)
	defer file.Close()
}

func main() {
	//basicoper()
	writefile()

}
