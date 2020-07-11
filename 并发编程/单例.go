package main

import (
	"fmt"
	"sync"
)

type User struct {
	Name string
}

func GetInstance() *User {
	new(sync.Once).Do(func() {
		user = &User{}
	})
	return user
}

var user *User

func inits() {
	//var user = User.GetInstance()

	var a = sync.WaitGroup{}
	for _ = range [10]int{} {
		a.Add(1)
		go func() {
			user = GetInstance()
			fmt.Printf("%p\n", &user)
			a.Done()
		}()
	}
	a.Wait()
}

func main() {
	inits()
}
