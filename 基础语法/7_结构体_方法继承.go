package main

import "fmt"

type person struct {
	id int
}

type male struct {
	person
	name string
}

func  (p *person)eat(){
	fmt.Println("吃东西的人是...",p.id)
}

func (m *male) eat(){
	fmt.Println("吃东西的学生是...",m.id)
}
func main() {
	ma := new(male)
	ma.id = 12345
	ma.eat()


}
