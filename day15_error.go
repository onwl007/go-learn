package main

import (
	"fmt"
	"os"
)

func main() {
	/**
	错误是什么？
		指的是可能出现问题的地方出现了问题。比如打开一个文件时失败，这种情况在人们的意料之中

	异常是什么？
		指的是不应该出现问题的出现了问题。比如引用了空指针，这种情况在人们对意料之外

	可见，错误是业务过程的一部分，而异常不是
	*/
	f, err := os.Open("/test.txt")
	if err, ok := err.(*os.PathError); ok {
		fmt.Println("File at path", err.Path, "failed to open")
		return
	}
	fmt.Println(f.Name(), "opened successfully")

	//如果一个函数或方法返回一个错误，那么按照惯例，必须是函数返回的最后一个值
	//常用方法是将返回的错误与 nil 进行比较。nil 值表示没有发生错误，而非 nil 值表示出现错误
}
