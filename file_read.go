package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/**
	读取数据
		Reader 接口：
			Read(p []byte)(n int, error)
	*/
	// step1: 打开文件
	fileName := "/Users/wanglei/go/src/go-learn/aa.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	// step3: 关闭文件
	defer file.Close()

	// step2：读取数据
	bs := make([]byte, 4, 4)
	// 第一次读取
	//n1, err1 := file.Read(bs)
	//fmt.Println(err1)
	//fmt.Println(n1)
	//fmt.Println(bs)
	//fmt.Println(string(bs))

	// 第二次读取
	//n2, err2 := file.Read(bs)
	//fmt.Println(err2)
	//fmt.Println(n2)
	//fmt.Println(bs)
	//fmt.Println(string(bs))

	// 第三次读取
	//n3, err3 := file.Read(bs)
	//fmt.Println(err3)
	//fmt.Println(n3)
	//fmt.Println(bs)
	//fmt.Println(string(bs))

	n := -1
	for {
		n, err = file.Read(bs)
		if n == 0 || err == io.EOF {
			fmt.Println("读取到了文件的末尾，结束读取操作。。。")
			break
		}
		fmt.Println(string(bs[:n]))
	}
}
