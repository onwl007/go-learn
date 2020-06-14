package main

import (
	"fmt"
	"math"
)

func main() {
	/**
	形参：定义函数时，用于接收外部传入的数据，叫做形式参数，简称形参
	实参：定义函数时，用于接收外部传入的数据，叫做实际参数，简称实参
	可变参：接受变参的函数是有着不定数量的参数的
	*/
	var a int = 100
	var b int = 200
	var ret int

	ret = max(a, b)
	fmt.Printf("最大值是：%d\n", ret)

	//值传递
	//声明函数变量
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	//使用函数
	fmt.Println(getSquareRoot(9))

	//引用传递 其实传递是变量的内存地址
	//变量在内存中是存放于一定地址上的，修改变量实际上就是修改变量地址处的内存
	/**
	1.传指针使得多个函数能操作统一对象
	2.传指针比较轻量级(8 bytes)，只是传内存地址，可以用指针传体积较大的结构体，会减少较多的系统开销
	3.slice，map这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针（注：若函数需要改变 slice 的长度，则仍需要取地址传递指针）
	*/

	x := 3
	fmt.Println("x = ", x)
	x1 := add1(&x)
	fmt.Println("x+1= ", x1)
	fmt.Println("x = ", x)

	// 一个函数可以没有返回值，也可以有一个返回值，也可以返回多个值
	//_ 是 Go 中的空白标识符，可以替代任何类型的返回值
	area, _ := rectProps(10.8, 5.6)
	fmt.Printf("Area %f \n", area)

	//函数的作用域
	//作用域：变量可以使用的范围
	//函数也是 Go 语言中的一种数据类型，可以作为另一个函数的参数，也可以作为另一个函数的返回值

	/**
	defer 函数
		延迟语句，被用于执行一个函数调用，在这个函数之前，延迟语句返回

		有多个 defer 语句时，函数执行到最后时，这些 defer 语句会按照逆序执行，最后该函数返回
		1.如果有很多调用 defer，那么 defer 是采用后进先出模式
		2.在离开所在的方法时，执行（报错的时候也会执行）

	延迟并不仅仅局限于函数，延迟一个方法调用也是完全合法的

	延迟参数
		延迟函数的参数在执行延迟语句时被执行，而不是在执行实际的函数调用时执行

	注意
		defer函数
		1.当外围函数中的语句正常执行完毕时，只有其中所有的延迟函数都执行完毕，外围函数才会真正的结束执行
		2.当执行外围函数中的 return 语句时，只有其中所有的延迟函数都执行完毕后，外围函数才会真正返回
		3.当外围函数中的代码引发运行恐慌时，只有其中所有的延迟函数都执行完毕后，该运行恐慌才会真正被扩展至调用函数
	*/
	m := 1
	n := 2
	defer fmt.Println(n)
	fmt.Println(m)

	nums := []int{78, 109, 2, 563, 300}
	largest(nums)

	p := person{
		firstName: "John",
		lastName:  "Smith",
	}
	defer p.fullName()
	fmt.Printf("Welcome ")

	a1 := 5
	defer printA(a1)
	a1 = 10
	fmt.Println("value of a before deferred function call", a1)

}

func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func add1(a *int) int {
	*a = *a + 1
	return *a
}

func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func finished() {
	fmt.Println("Finished finding largest")
}

func largest(nums []int) {
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
}

type person struct {
	firstName string
	lastName  string
}

func (p person) fullName() {
	fmt.Printf("%s %s \n", p.firstName, p.lastName)
}

func printA(a int) {
	fmt.Println("value of a in deferred function", a)
}
