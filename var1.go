package main

import "fmt"

/**
变量：variable
概念：一小块内存，用于存储数据，在程序运行过程中数值可以改变\

使用：
	step1：变量的声明，也叫定义
			第一种：var 变量名 数据类型
					变量名 = 赋值
					var 变量名 数据类型 = 赋值
			第二种：类型推断，省略 var
					变量名 := 赋值
	step2：变量的访问，赋值和取值
			直接变量名访问
*/

func main() {
	// 第一种：定义变量，然后进行赋值
	var num1 int
	num1 = 30
	fmt.Printf("num1的数值是：%d\n", num1)
	// 写在一行
	var num2 int = 15
	fmt.Printf("num2的数值是：%d\n", num2)

	// 第二种：类型推断
	var name = "张祎"
	fmt.Printf("类型是：%T，数值是：%s\n", name, name)

	// 第三种，简短定义，也叫简短声明
	sum := 100
	fmt.Println(sum)

	// 多个变量同时声明
	var a, b, c int
	a = 1
	b = 2
	c = 3
	fmt.Println(a, b, c)

	var m, n int = 100, 200
	fmt.Println(m, n)

	var n1, f1, s1 = 100, 3.14, "Go"
	fmt.Println(n1, f1, s1)

	var (
		studentName = "张祎"
		age         = 26
		sex         = "女"
	)
	fmt.Printf("学生姓名：%s, 年龄：%d，性别：%s", studentName, age, sex)
}
