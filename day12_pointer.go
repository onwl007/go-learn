package main

import "fmt"

type name int8
type first struct {
	a int
	b bool
	name
}

const MAX int = 3

func main() {
	/**
	指针
		存储另一个变量的内存地址的变量

	声明指针，*T 是指针变量的类型，它指向 T 类型的值，* 号用于指定变量是作为一个指针
	当一个指针被定义后没有分配到任何变量时，它的值为 nil
	nil 指针也称为空指针

	获取指针的值
		访问指针指向的变量的值，语法 *a

	Go 不支持指针算法
		b:=2
		p:=&b
		p++ // invalid operation: p++ (non-numeric type *int)

	指针的指针
		如果一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量
		var ptr **int
	*/

	var a int = 20 //声明实际变量
	var ip *int    //声明指针变量

	ip = &a //指针变量的存储地址
	fmt.Printf("a 变量的地址：%x\n", &a)

	//指针变量的存储地址
	fmt.Printf("ip 变量的存储地址：%x\n", ip)

	//使用指针访问值
	fmt.Printf("*ip 变量的值：%d\n", *ip)

	b1 := 255
	a1 := &b1
	fmt.Println("address of b1 is ", a1)
	fmt.Println("value of b1 is ", *a1)
	*a1++
	fmt.Println("new value of b1 is ", b1)

	a2 := 58
	fmt.Println("value of a2 before function call is ", a2)
	b2 := &a2
	change(b2)
	fmt.Println("value of a2 after function call is ", a2)

	//不要将一个指向数组的指针传递给函数，使用切片
	//虽然将指针传递给一个数组作为函数的参数并对其进行修改，但这并不是实现这一目标的惯用方法，我们有切片
	a3 := [3]int{89, 90, 91}
	//modify(&a3)
	fmt.Println(a3[:])
	modyifSlice(a3[:])
	fmt.Println(a3)

	a4 := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a4[i]
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a4[%d] =%d\n", i, *ptr[i])
	}

	// 指针的指针
	var a5 int
	var ptr5 *int
	var pptr5 **int

	a5 = 3000

	//指针 ptr5 地址
	ptr5 = &a5

	//指向指针 pptr5 地址
	pptr5 = &ptr5

	//获取 pptr5 的值
	fmt.Printf("变量 a5 = %d\n", a5)
	fmt.Printf("指针变量 *ptr5 = %d\n", *ptr5)
	fmt.Printf("指向指针的指针变量 **ptr5 = %d\n", **pptr5)

	var a6 int = 100
	var b6 int = 200

	fmt.Printf("交换前 a6 的值：%d\n", a6)
	fmt.Printf("交换前 b6 的值：%d\n", b6)

	//调用函数用于交换值
	//&a6 指向 a6 变量的地址
	//&b6 指向 b6 变量的地址
	swap(&a6, &b6)

	fmt.Printf("交换后 a6 的值：%d\n", a6)
	fmt.Printf("交换后 b6 的值：%d\n", b6)

}

func swap(x *int, y *int) {
	var temp int
	temp = *x //保存 x 地址的值
	*x = *y   //将 y 赋值给 x
	*y = temp //将 temp 赋值给 y
}

func change(val *int) {
	*val = 55
}

func modify(arr *[3]int) {
	(*arr)[0] = 90
}

func modyifSlice(sls []int) {
	sls[0] = 90
}
