package main

import "fmt"

type add interface {
	result()
}
type Person struct {
	name string
}

func (per *Person) result()  {
	fmt.Println(per.name)
}

func Bian(a add)  {
	a.result()
}
func main() {

	per := Person{"cao"}
	fmt.Println(per.name)

	ad := per
	ad.result()
	fmt.Printf("%p",&ad)
	fmt.Println("")
	fmt.Printf("%p",&per)

	Bian(&per)
}
