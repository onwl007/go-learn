package main

import (
	"fmt"
	"math"
)

type Employee struct {
	name     string
	salary   int
	currency string
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

type HumanKind struct {
	name  string
	age   int
	phone string
}

type StudentKind struct {
	HumanKind
	school string
}

type EmployeeKind struct {
	HumanKind //匿名字段
	company   string
}

func (h *HumanKind) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//Employee 的 method 重写 HumanKind 的 method
func (e *EmployeeKind)SayHi(){
	fmt.Printf("Hi, I am %s, I work at %s. call me on %s\n",e.name,e.company,e.phone)
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

//可用重名，该 method 属于 Circle 类型对象中的方法
func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

func main() {
	/**
	方法
		一个方法就是包含了接受者的函数，可以是命名类型或者结构体类型的一个值或者是一个指针

	和函数的区别
		方法有一个接收者，给一个函数加一个接收者，就变成了方法，接收者可以是值接收者，也可以是指针接收者

	继承
		如果匿名字段实现了一个 method，那么包含这个匿名字段的 struct 也能调用该 method
		方法是可以继承和重写的
		存在继承关系时，按照就近原则，进行调用
	*/
	emp1 := Employee{name: "Sam Adolf", salary: 5000, currency: "$"}
	emp1.displaySalary()

	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}
	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())

	mark := StudentKind{HumanKind{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := EmployeeKind{HumanKind{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}
	mark.SayHi()
	sam.SayHi()
}
