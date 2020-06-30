package main

import "fmt"

func main() {

	//arr()
	//slice()
	//rangem()
	//mmap()
	//typecast()
	arrtest()
}

func arrtest() {
	a := []int{1, 2, 3, 4}
	fmt.Println(a[1:3])
}

func typecast() {
	var age int = 17
	var firstname = "cao"
	var height float64 = 175.2
	a := float32(height)
	fmt.Println(age, firstname, height, a)
	b := height - float64(17)
	fmt.Println(b)
}

func mmap() {
	var countrycenter = make(map[string]string)
	countrycenter["shanxi"] = "taiyuan"
	countrycenter["liaoning"] = "shenyang"
	countrycenter["jilin"] = "changchun"

	for country, center := range countrycenter {
		fmt.Println(country, center)
	}

	//删除元素
	delete(countrycenter, "shanxi")
	fmt.Println(countrycenter)

}

func rangem() {
	s := make([]int, 5, 10)
	for i, v := range s {
		fmt.Println(i, v)
	}
	b := true
	a := false
	if a && b {
		fmt.Println(b)
	}
	//a  := (2 && 1)
	//fmt.Println(a)
}

func slice() {
	var slice []int
	println(slice)
	var slice1 = []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)
	var slice2 = make([]int, 2, 6)
	fmt.Println(slice2)

	fmt.Println(cap(slice2), slice1[2:4])
	slice2 = append(slice2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6)
	fmt.Println(cap(slice2))
	fmt.Println(slice2)

	s := make([]int, 5, 8)
	fmt.Printf("%p\n", s)
	s = append(s, 5, 6, 7)
	fmt.Printf("%p\n", s)

	a := s[0:5:5]
	fmt.Println(a)

	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(s1[2:3:3])

	copy(s1, s)
	fmt.Println(s1, s)

}

func arr() {
	arr := [...]int{1, 2, 3, 4, 5}

	s := arr[1:]
	for i := 0; i < len(s); i++ {
		println(s[i])
	}
	x := make([]int, 3, 5)
	for i := 0; i < len(x); i++ {
		println(x[i])
	}
}
