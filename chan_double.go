package main

import "fmt"

func main() {
	/**
	双向：
		chan T
		chan <- data, 发送数据，写出
		data <- chan，获取数据，读取

	单向：定向
		chan <- T, 只支持写
		<- chan T, 只读
	*/
	ch1 := make(chan string)
	done := make(chan bool)
	go sendData(ch1, done)

	data := <-ch1
	fmt.Println("子goroutine传来，", data)

	ch1 <- "我是main"

	<-done
	fmt.Println("main...over...")
}

func sendData(ch1 chan string, done chan bool) {
	ch1 <- "我是你爹"
	data := <-ch1
	fmt.Println("main goroutine 传来，", data)

	done <- true
}
