package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//获取goroot目录
	fmt.Println("GOROOT--->", runtime.GOROOT())

	//获取操作系统
	fmt.Println("os/platform--->", runtime.GOOS)

	//获取逻辑cpu数量
	fmt.Println("逻辑cpu的数量--->", runtime.NumCPU())

	//go func() {
	//	for i := 0; i < 5; i++ {
	//		fmt.Println("goroutine...")
	//	}
	//}()
	//
	//for i := 0; i < 4; i++ {
	//	//让出时间片，先让别的goroutine执行
	//	runtime.Gosched()
	//	fmt.Println("main...")
	//}

	go func() {
		fmt.Println("goroutine开始...")
		fun()
		fmt.Println("goroutine结束...")
	}()

	//睡一会
	time.Sleep(3 * time.Second)
}

func fun() {
	defer fmt.Println("derfer...")
	//return //终止函数
	runtime.Goexit() //终止当前的goroutine
	fmt.Println("fun函数...")
}
