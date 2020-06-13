package main

import "fmt"

func main() {
	/**
	位运算符：
		将数值，转为二进制后，按位操作
	按位&：
		对应位的值如果都为1才为1，有一个为0就位0
	按位|：
		对应位的值如果都是0才是0，有一个为1就为1
	异或^：
		二元：a^b
			对应位的值不同为1，相同为0
		一元：^a
			按位取反
				1---->0
				0---->1
		位清空：&^
			对于 a &^ b
			如果为0，则取a对应位上的数值
			如果为1，则结果位就取0

	位移运算符：
	<<：按位左移，将a转为二进制，向左移动b位
	>>：按位右移，将a转为二进制，向右移动b位
	*/

	a := 60
	b := 13
	/**
	a:60 0011 1100
	b:13 0000 1101
	&:12 0000 1100
	|:   0011 1101
	*/
	fmt.Printf("a: %d, %b\n", a, a)
	fmt.Printf("b: %d, %b\n", b, b)

	res1 := a & b
	fmt.Println(res1)

	res2 := a | b
	fmt.Println(res2)

	res3 := a ^ b
	fmt.Println(res3)

	res4 := a &^ b
	fmt.Println(res4)

	res5 := ^a
	fmt.Println(res5)

	c := 8
	res6 := c << 2
	fmt.Println(res6)

	res7 := c >> 2
	fmt.Println(res7)
}
