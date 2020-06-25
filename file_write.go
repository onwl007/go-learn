package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	/**
	写数据
	*/

	fileName := "/Users/wanglei/go/src/go-learn/aa.txt"
	//step1:打开文件
	//step1:写出数据
	//step1:关闭文件

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// 写出数据
	bs := []byte{97, 98, 99, 100} // a,b,c,d
	n, err := file.Write(bs[:2])
	fmt.Println(n)
	handleErr(err)
	file.WriteString("\n")

	// 直接写出字符串
	n, err = file.WriteString("helloworld")
	fmt.Println(n)
	handleErr(err)

	file.WriteString("\n")
	n, err = file.Write([]byte("today"))
	fmt.Println(n)
	handleErr(err)
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
