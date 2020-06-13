package main

import "fmt"

func main() {
	var a int = 8
	if a < 20 {
		fmt.Printf("a 小于 20\n")
	}
	fmt.Printf("a 的值为：%d\n", a)

	if num := 10; num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "D" // case 后可以有多个数值
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Println("优秀")
	case grade == "B", grade == "C":
		fmt.Println("良好")
	case grade == "D":
		fmt.Println("及格")
	case grade == "F":
		fmt.Println("不及格")
	default:
		fmt.Println("差")
	}
	fmt.Printf("你的等级是 %s\n", grade)

	type data [2]int
	switch x := 5; x {
	default:
		fmt.Println(x)
	case 5:
		x += 10
		fmt.Println(x)
		fallthrough
	case 6:
		x += 20
		fmt.Println(x)

	}

	num1 := 75
	switch {
	case num1 >= 0 && num1 <= 50:
		fmt.Println("num is greater than 0 and less than 50")
	case num1 >= 51 && num1 <= 100:
		fmt.Println("num is greater than 51 and less than 100")
	case num1 >= 101:
		fmt.Println("num is greater than 100")
	}

	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型：%T\n", i)
	case int:
		fmt.Println("x 是 int 型")
	case float64:
		fmt.Println("x 是 float64 型")
	case func(int) float64:
		fmt.Println("x 是 func(int) 型")
	case bool, string:
		fmt.Println("x 是 bool 或 string 型")
	default:
		fmt.Println("未知型")

	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d", i)
	}

	var b int = 15
	var a1 int

	numbers := [6]int{1, 2, 3, 5}
	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值为：%d\n", a)
	}

	for a1 < b {
		a1++
		fmt.Printf("a1 的值为：%d\n", a1)
	}

	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

	for i := 1; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Printf("%d", i)
	}
	fmt.Printf("\nline after for loop\n")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d\n", i)
	}

	var a2 int = 10
LOOP:
	for a2 < 20 {
		if a2 == 15 {
			// 跳过迭代
			a2 = a2 + 1
			goto LOOP
		}
		fmt.Printf("a2的值为：%d\n", a2)
		a2++
	}

	var arr [4]float32
	fmt.Println(arr)

	var arr1 = [5]string{"ruby", "王二狗", "rose"}
	fmt.Println(arr1)

	var arr2 = [5]int{'A', 'B', 'C', 'D', 'E'}
	fmt.Println(arr2)

	arr3 := [...]int{1, 2, 3, 4, 5} // 根据元素个数，设置数组的大小
	fmt.Println(arr3)

	arr4 := [5]int{4: 100} // [0 0 0 0 100]
	fmt.Println(arr4)

	arr5 := [...]int{0: 1, 4: 1, 9: 1} // [1 0 0 0 1 0 0 0 0 1]
	fmt.Println(arr5)

	var arr6 [10]int

	for i := 0; i < 10; i++ {
		arr6[i] = i + 100
	}

	for j := 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, arr6[j])
	}

	arr7 := [...]float64{67.7, 89.8, 21, 78}
	fmt.Println("length of a is", len(arr7))

	arr8 := [...]int{12, 78, 50}
	fmt.Println(arr8)

	for i := 0; i < len(arr7); i++ {
		fmt.Printf("%d th element of a is %.2f\n", i, arr7[i])
	}

	sum := float64(0)
	for i, v := range arr7 {
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a", sum)

	for _, v := range arr7 {
		fmt.Printf("the element of a is %.2f\n", v)
	}

	arr9 := [3][4]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
	}
	fmt.Println(arr9)

	arr10 := [...]string{"USA", "China", "India", "Germany", "France"}
	arr11 := arr10
	arr11[0] = "Singapore"
	fmt.Println("arr10 is ", arr10)
	fmt.Println("arr11 is ", arr11)

	// 数组的大小也是类型的一部分，因此数组不能被调整大小，不要担心这个限制，因为切片的存在就是为了解决这个问题
	//arr12 := [3]int{5, 78, 8}
	//var arr13 [5]int
	//arr13 = arr12 // cannot use arr12 (type [3]int) as type [5]int in assignment

	/**
	切片
		切片是一种方便、灵活且强大的包装器，本身没有任何数据，只是对现有数组的引用
		切片更像是一个结构体，包含了三个元素
			1.指针，指向数组中slice指定的开始位置
			2.长度，即slice的长度
			3.最大长度，也就是slice开始位置到数组的最后位置的长度
	*/
	arr12 := [5]int{76, 77, 78, 79, 80}
	var arr13 []int = arr12[1:4]
	fmt.Println(arr13)

	// slice 没有自己的任何数据，只是底层数组的一个展示，对 slice 所做的任何修改都将反映在底层数组中
	arr14 := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	arr14Slice := arr14[2:5]
	fmt.Println("array before", arr14)
	for i := range arr14Slice {
		arr14Slice[i]++
	}
	fmt.Println("array after", arr14)

	numa := [3]int{78, 79, 80}
	nums1 := numa[:]
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)

	var slice1 = make([]int, 3, 5)
	printSlice(slice1)

	var slice2 []int
	printSlice(slice2)
	if slice2 == nil {
		fmt.Println("切片是空的")
	}

	slice3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	printSlice(slice3)

	//打印原始切片
	fmt.Println("slice3 == ", slice3)

	//打印子切片从索引1到索引4（不包含）
	fmt.Println("slice3[1:4] == ", slice3[1:4])

	//默认下限为0
	fmt.Println("slice3[:3] == ", slice3[:3])

	//默认上限为 len(s)
	fmt.Println("slice3[4:] == ", slice3[4:])

	slice4 := make([]int, 0, 5)
	printSlice(slice4)

	//打印子切片从索引0到索引2
	slice5 := slice3[:2]
	printSlice(slice5)

	//打印子切片从索引2到索引5
	slice6 := slice3[2:5]
	printSlice(slice6)

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
