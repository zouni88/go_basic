package main

import (
	"fmt"
	"runtime"
)

type Vector []float64

var CpuNo = 8

func gen() []float64 {
	var a = make([]float64, 100000000, 100000000)
	for i, _ := range [100000000]int{} {
		a[i] = float64(i)
	}
	return a
}

func Compute(unit []float64, c chan int) {
	sum := 0.
	for _, v := range unit {
		sum += v
	}
	fmt.Println(sum)
	c <- 1
}

func main() {
	//应用多CPU进行并行任务
	runtime.GOMAXPROCS(1)
	a := gen()
	c := make(chan int)
	setup := len(a) / 8
	fmt.Println("setop = ", setup)
	start := 0
	stop := setup
	for i := 0; i < CpuNo; i++ {
		unit := a[start:stop]
		start += setup
		stop = start + setup
		go Compute(unit, c)
	}

	for i := 0; i < CpuNo; i++ {
		res := <-c
		fmt.Printf("结束一个%v  %v \n", i, res)
	}

	fmt.Println("结束")
}
