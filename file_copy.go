package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 在性能上，不管是io.Copy()还是ioutil包，性能上都是还不错的
func main() {
	srcFile := "/Users/wanglei/Downloads/极客时间算法训练营/第三课/2. 实战题目解析：移动零.mov"
	destFile := "/Users/wanglei/go/src/go-learn/test.mov"

	//total, err := copyFile1(srcFile, destFile)
	//total, err := copyFile2(srcFile, destFile)
	total, err := copyFile3(srcFile, destFile)
	fmt.Println(err)
	fmt.Println(total)
}

func copyFile1(srcFile, destFile string) (int, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}

	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}

	defer file1.Close()
	defer file2.Close()

	// 拷贝数据
	bs := make([]byte, 1024, 1024)
	n := -1 // 读取的数据量
	total := 0
	for {
		n, err = file1.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("拷贝完毕")
			break
		} else if err != nil {
			fmt.Println("报错了。。。")
			return total, nil
		}
		total += n
		file2.Write(bs[:n])
	}
	return total, nil
}

func copyFile2(srcFile, destFile string) (int64, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}

	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}

	defer file1.Close()
	defer file2.Close()

	return io.Copy(file2, file1)
}

// 由于使用一次性读取文件，再一次性写入文件的方式，所以该方法不适用于大文件，容易内存溢出
func copyFile3(srcFile, destFile string) (int, error) {
	input, err := ioutil.ReadFile(srcFile)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = ioutil.WriteFile(destFile, input, 0644)
	if err != nil {
		fmt.Println("操作失败：", destFile)
		fmt.Println(err)
		return 0, err
	}

	return len(input), nil
}
