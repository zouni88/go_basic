package models

import "fmt"

type Huawei struct {
	Name string

}

func (huawei *Huawei) Call() {
	fmt.Printf("这是一个%s手机\n", huawei.Name)
}
