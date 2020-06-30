package main

import (
	"go_basic/models"
)

func Call(phone models.Phone) {
	phone.Call()
}

func main() {
	hw := models.Huawei{Name: "P30"}
	iphone := models.Iphone{Name: "iphone 11"}
	Call(&hw)
	Call(&iphone)

}
