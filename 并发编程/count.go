package main

import (
	"fmt"
	"time"
)

type Person struct {
	Money int
}

func A(msgs chan *Person, name string) {
	per := <-msgs
	per.Money--
	fmt.Printf("name is %v money==%v\n", name, per.Money)
	//pe := *per
	msgs <- per

	//runtime.GOMAXPROCS()
}

func delay(msgs chan *Person, per *Person) {
	time.Sleep(time.Second * 2)
	fmt.Println("。。。")
	msgs <- per
	fmt.Println("。。。")
	fmt.Println(msgs)

}

func main() {
	person := Person{Money: 10}
	msgs := make(chan *Person)
	go A(msgs, "A")
	go A(msgs, "B")
	delay(msgs, &person)
	msgs <- &person
	fmt.Println(person.Money)
}
