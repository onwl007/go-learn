package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // 创建同步等待组的对象

func main() {
	/**
	WaitGroup：同步等待
		Add(), 设置等待组中要执行的子 goroutine 的数量

		Wait(), 让主 goroutine 出于等待
	*/
	wg.Add(2)
	go fun1()
	go fun2()

	fmt.Println("main 进入阻塞状态。。等待wg中的子goroutine结束。。")
	wg.Wait() // 表示 main goroutine 进入等待，意味着阻塞
	fmt.Println("main...解除阻塞。。")
}

func fun1() {
	for i := 1; i < 10; i++ {
		fmt.Println("fun1()函数打印。。。A ", i)
	}

	wg.Done() // 给wg等待组中的counter的数值减1。同 wg.Add(-1)
}

func fun2() {
	defer wg.Done()
	for j := 1; j < 10; j++ {
		fmt.Println("\tfun2()函数打印。。。", j)
	}
}
