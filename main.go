package main

import (
	"fmt"
	"go_basic/models"

)

func main() {
	var hw models.Phone = &models.Huawei{Name: "P30"}
	var iphone models.Phone = &models.Iphone{Name: "iphone11"}
	iphone.Call()
	hw.Call()
	go func() {
		fmt.Println(123)
	}()

	var json struct{
		Name string `json:"name"`
	}

	json.Name = "cao"
	fmt.Println(json.Name)

}
