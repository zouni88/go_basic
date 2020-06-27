package main

import (
	"fmt"
	"sync"
)

type User struct {
	Name string
}

func GetInstance() * User {
	new(sync.Once).Do(func() {
		user = &User{}
	})
	return user
}

var user *User

func inits() {
	for _ = range [10]int{} {
		user = GetInstance()
		fmt.Printf("%p\n",&user)
	}
}

func main() {
	inits()
}
