package main

import "fmt"

func main() {
	name := "root"
	password := 1234565
	conn := fmt.Sprintf("%v:%v@tcp(zouni.vip:3306)/beego?charset=utf8", name, password)
	fmt.Println(conn)
}
