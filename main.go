package main

import (
	"go_basic/models"

)

func main() {
	var hw models.Phone = &models.Huawei{Name: "P30"}
	var iphone models.Phone = &models.Iphone{Name: "iphone11"}
	iphone.Call()
	hw.Call()
	go func() {
		print(123)
	}()

}
