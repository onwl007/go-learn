package main

import (
	"fmt"
	"os"
)

func main() {
	/**
	文件操作
	1.路径
		相对路径：relative
			aa.txt
		绝对路径：absolute
			/Users/ruby/Documents/pro/a/aa.txt

		.当前目录
		..上一层
	2.创建文件夹，如果文件夹存在，创建失败
		os.MkDir() 创建一层
		os.MkDirAll() 创建多层
	3.创建文件，Create 采用模式 0666 （任何人都可读写，不可执行）创建一个名为 name 的文件，如果文件已存在会截断它（为空文件）
		os.Create() 创建文件
	4.打开文件：让当前的程序和指定的文件之间建立一个连接
		os.Open(filename)
		os.OpenFile(filename,mode,perm)
	5.关闭文件：程序和文件之间的连接断开
		file.Close()
	6.删除文件或目录：慎用
		os.Remove() 删除文件和空目录
		os.RemoveAll() 删除所有

	*/
	//fileInfo, err := os.Stat("./aa.txt")
	//if err != nil {
	//	fmt.Println("err: ", err)
	//	return
	//}
	//fmt.Printf("%T\n", fileInfo)
	//// 文件名
	//fmt.Println(fileInfo.Name())
	//// 文件大小
	//fmt.Println(fileInfo.Size())
	//// 是否是目录
	//fmt.Println(fileInfo.IsDir())
	//// 修改时间
	//fmt.Println(fileInfo.ModTime())
	//// 权限
	//fmt.Println(fileInfo.Mode())

	// 创建目录
	//err := os.Mkdir("/Users/wanglei/go/src/go-learn/test", os.ModePerm)
	//if err != nil {
	//	fmt.Println("err: ", err)
	//	return
	//}
	//fmt.Println("文件夹创建成功")

	//err := os.MkdirAll("/Users/wanglei/go/src/go-learn/hello/world", os.ModePerm)
	//if err != nil {
	//	fmt.Println("err: ", err)
	//	return
	//}
	//fmt.Println("多层文件夹创建成功")

	// 创建文件
	//file, err := os.Create("/Users/wanglei/go/src/go-learn/hello.txt")
	//if err != nil {
	//	fmt.Println("err: ", err)
	//	return
	//}
	//fmt.Println(file)

	err := os.RemoveAll("/Users/wanglei/go/src/go-learn/hello")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("删除目录成功")
}
