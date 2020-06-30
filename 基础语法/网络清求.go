package main

import (
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.baidu.com")

	fmt.Println(res.Status)
	fmt.Println(err)

}
