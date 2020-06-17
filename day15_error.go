package main

import (
	"fmt"
	//"os"
	//"net"
	//"path/filepath"
	"math"
	//"errors"
)

type areaError struct {
	err    string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		//return 0, errors.New("Area calculation failed,radius is less than zero")
		//return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
		return 0, &areaError{"radius is negative", radius}
	}
	return math.Pi * radius * radius, nil
}

type rectangleError struct {
	err    string
	length float64
	width  float64
}

func (e *rectangleError) Error() string {
	return e.err
}

func (e *rectangleError) lengthNegative() bool {
	return e.length < 0
}

func (e *rectangleError) widthNegative() bool {
	return e.width < 0
}

func rectArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += ", width is less than zero"
		}
	}
	if err != "" {
		return 0, &rectangleError{err, length, width}
	}
	return length * width, nil
}

func main() {
	/**
	错误是什么？
		指的是可能出现问题的地方出现了问题。比如打开一个文件时失败，这种情况在人们的意料之中

	异常是什么？
		指的是不应该出现问题的出现了问题。比如引用了空指针，这种情况在人们对意料之外

	可见，错误是业务过程的一部分，而异常不是
	*/

	//f, err := os.Open("/test.txt")
	//if err, ok := err.(*os.PathError); ok {
	//	fmt.Println("File at path", err.Path, "failed to open")
	//	return
	//}
	//fmt.Println(f.Name(), "opened successfully")

	//如果一个函数或方法返回一个错误，那么按照惯例，必须是函数返回的最后一个值
	//常用方法是将返回的错误与 nil 进行比较。nil 值表示没有发生错误，而非 nil 值表示出现错误

	//addr, err1 := net.LookupHost("golangbot123.com")
	//if err1, ok := err1.(*net.DNSError); ok {
	//	if err1.Timeout() {
	//		fmt.Println("operation timed out")
	//	} else if err1.Temporary() {
	//		fmt.Println("temporary error")
	//	} else {
	//		fmt.Println("generic error: ", err1)
	//	}
	//	return
	//}
	//fmt.Println(addr)

	//files, error := filepath.Glob("[")
	//if error != nil && error == filepath.ErrBadPattern {
	//	fmt.Println(error)
	//	return
	//}
	//fmt.Println("matched files", files)

	//files, _ := filepath.Glob("[")
	//fmt.Println("matched files", files)

	//radius := -20.0
	//area, err := circleArea(radius)
	//if err != nil {
	//	if err, ok := err.(*areaError); ok {
	//		fmt.Printf("Radius %0.2f is less than zero", err.radius)
	//		return
	//	}
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Printf("Area of circle %0.2f\n", area)

	length, width := -5.0, -9.0
	rect, err := rectArea(length, width)
	if err != nil {
		if err, ok := err.(*rectangleError); ok {
			if err.lengthNegative() {
				fmt.Printf("error: length %0.2f is less than zero\n", err.length)
			}
			if err.widthNegative() {
				fmt.Printf("error: width %0.2f is less than zero\n", err.width)
			}
		}
		fmt.Println(err)
		return
	}
	fmt.Println("area of rect", rect)
}
