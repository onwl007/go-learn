package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	/**
	Go 中，接口是一组方法签名，当类型为接口中的所有方法提供定义时，被称为实现接口
	1.它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口
	2.接口定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就实现了该接口
	3.interface 可以被任意的对象实现
	4.任意的类型都实现了空 interface ，也就是包含 0 个 method 的 interface

	interface 可以存储实现了这个 interface 的任意类型的对象
	*/
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
