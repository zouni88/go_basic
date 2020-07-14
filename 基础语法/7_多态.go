package main

import (
	"go_basic/models"
)

func Call(phone models.Phone) {
	phone.Call()
}

func main() {
	var hw models.Phone = &models.Huawei{Name: "P30"}
	var hh models.Phone = new(models.Huawei)
	hh.Call()
	iphone := models.Iphone{Name: "iphone 11"}
	Call(hw)
	Call(&iphone)

}
