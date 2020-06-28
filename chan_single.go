package main

import (
	"debug/elf"
	"fmt"
)

func main() {
	/**
	单向：定向
		chan <- T 只支持写
		<- chan T 只读
	*/
	ch1 := make(chan int) // 双向，读，写
	//ch2 := make(chan<- int) // 单向，只能写，不能读
	//ch3 := make(<-chan int) // 单向，只能读，不能写

	//ch1 <- 100
	//data := ch1

	//ch2 <- 1000
	//data := <-ch2

	//data := <-ch3
	//ch3 <- 200

	go fun1(ch1) //可读，可写
	//go fun1(ch2) // 只写

	data := <-ch1
	fmt.Println("fun1函数中写出的数据是：", data)

	go fun2(ch1)
	ch1 <- 200
	fmt.Println("mian...over...")
}

//该函数，只能操作只写的通道
func fun1(ch chan<- int) {
	ch <- 100
	fmt.Println("fun1函数结束。。。")
}

//该函数，只能操作只读的通道
func fun2(ch <-chan int) {
	data := <-ch
	fmt.Println("fun2函数，从ch中读取的通道是：", data)
}
