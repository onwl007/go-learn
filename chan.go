package main

import "fmt"

func main() {
	/**
	channel 通道
	*/
	var a chan int
	fmt.Printf("%T, %v\n", a, a)

	if a == nil {
		fmt.Println("channel 是 nil 的，不能使用，需要先创建通道")
		a = make(chan int)
		fmt.Println(a)
	}
}

func test1(ch chan int) {
	fmt.Printf("%T, %v\n", ch, ch)
}
