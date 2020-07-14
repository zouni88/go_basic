package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	var a *int
	b := 123
	a = &b
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(&b)

	fmt.Println(*a)

	aa := 1
	bb := 2
	swap(&aa, &bb)
	fmt.Println(aa, bb)

	cc := "cao"
	fmt.Printf("%p\n",&cc)
	ee := &cc

	*ee = "zhang"
	fmt.Printf("%p\n",&cc)

	var eee = make([]byte,1)
	var bbb byte= 'c'
	eeee := byte('0')
	var ddd byte = 'd'
	eee = append(eee,ddd)
	fmt.Println(string(bbb),"\n",bbb)
	eee = append(eee,ddd,eeee)
	fmt.Println(eee)

	//ddddd := []byte{ddd}
	//eee = append(eee,ddddd)
	//append

}
