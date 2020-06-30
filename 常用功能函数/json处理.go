package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Game struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	game := Game{Id: 1, Name: "望着联盟"}
	res, _ := json.Marshal(game)

	fmt.Println(reflect.TypeOf(res))
	z := string(res)
	fmt.Println(z)
}
