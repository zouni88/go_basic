package main

import "fmt"

type add interface {
	result()
}
type Person struct {
	name string
}

func (per *Person) result() {
	fmt.Println(per.name)
}

func Bian(a add) {
	a.result()
}

func modifyParam(count *int,modify *ParamModify)  {
	*count = 123
	modify.Name = "caowang"
}

type ParamModify struct {
	Name string
}

func main() {

	per := Person{"cao"}
	fmt.Println(per.name)

	ad := per
	ad.result()
	fmt.Printf("%p", &ad)
	fmt.Println("")
	fmt.Printf("%p", &per)

	Bian(&per)

	var count = 111
	var modify = ParamModify{Name: "zhang"}
	modifyParam(&count,&modify)
	println(count,modify.Name)

	var aaa = 123
	var bbb *int
	bbb = &aaa
	*bbb = 666
	println(aaa)

}
