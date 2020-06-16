package main

import (
	"fmt"
	"strconv"
)

type my_fun func(int, int) string

func fun1() my_fun {
	fun := func(a, b int) string {
		s := strconv.Itoa(a) + strconv.Itoa(b)
		return s
	}

	return fun
}

type myint int    // 跟 int 相比，相当于两个不同的类型
type myint2 = int //不是重新定义类型，只是给int起别名

type Person1 struct {
	name string
}

func (p Person1) Show() {
	fmt.Println("Person-->", p.name)
}

//类型别名
type People1 = Person1

type Student1 struct {
	//嵌入两个结构
	Person1
	People1
}

func (p People1) Show2() {
	fmt.Println("People----->", p.name)
}

func main() {
	/**
	支持函数式编程，函数可以作为另一个函数的参数，也可以作为另一个函数的返回值

	非本地类型不能定义方法
		指不能为不在一个包中的类型定义方法
		解决方法
			1.将类型别名改为类型定义
			2.将别名定义放在相同包中
	*/
	res1 := fun1()
	fmt.Println(res1(10, 20))

	var i1 myint
	var i2 = 100
	i1 = 100
	fmt.Println(i1)
	fmt.Println(i1, i2)

	var i3 myint2
	i3 = i2
	fmt.Println(i1, i2, i3)

	//在结构体成员嵌入时使用别名
	//在通过 s 直接访问 name 的时候，或者 s 直接调用 Show() 方法，因为两个类型
	//都有 name 字段和 Show() 方法，会发生歧义，证明 People 的本质确实是 Person 类型
	var s Student1
	//s.name = "王二狗" // ambiguous selector s.name
	s.People1.name = "李小花"
	s.Person1.name = "王二狗"
	//s.Show() //ambiguous selector s.Show
	s.Person1.Show()
	s.People1.Show2()
	fmt.Printf("%T,%T\n,", s.Person1, s.People1)
}
