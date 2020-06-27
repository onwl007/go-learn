package main

import (
	"fmt"
	"time"
)

func main() {
	/**
	一个goroutine打印数字，另外一个goroutine打印字母，观察运行结果
	*/

	//1.先创建并启动子goroutine，执行printNum()
	go printNum()

	//2.main中打印字母
	for i := 1; i <= 1000; i++ {
		fmt.Printf("\t主goroutine中打印字母：A %d\n", i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("main...over...")
}

func printNum() {
	for i := 1; i <= 1000; i++ {
		fmt.Printf("子goroutine中打印数字：%d\n", i)
	}
}
