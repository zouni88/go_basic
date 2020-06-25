package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func writeFile(){
	file ,_ := os.OpenFile("aa.txt",os.O_RDWR,os.FileMode(6))
	file.Write([]byte("我的天1\r\n"))
	file.Write([]byte("我的天2\n\n"))
	defer file.Close()

}

func buffoper() {
	file,_ := os.Open("aa.txt")
	reader := bufio.NewReader(file)

	for {
		res,err := reader.ReadBytes('\n')
		if err == io.EOF{
			fmt.Println("结束了")
			break
		}
		fmt.Println(string(res))
	}

	file.Close()

}

func main() {
	//writeFile()
	buffoper()
}
