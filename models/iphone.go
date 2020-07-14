package models

import "fmt"


type Iphone struct {
	Name string
}

func (iphone *Iphone) Call() {
	fmt.Printf("这是一个%s手机\n", iphone.Name)
}