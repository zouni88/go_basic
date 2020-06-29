package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(time.Second)
	for {
		times := <- t.C
		fmt.Println("...",times)
	}

}
