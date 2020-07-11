package main

import (
	"fmt"
	"net/http"
)

func main() {
	res,err := http.Get("http://www.zouni.vip")

	fmt.Println(res)
	fmt.Println(err)
}
