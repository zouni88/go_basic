package models

import "fmt"

func (iphone *Iphone) Call()  {
	fmt.Printf("这是一个%s手机",iphone.Name)
}

type Iphone struct {
	Name string
}
