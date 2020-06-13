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
}
