package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	dirname := "/Users/wanglei/workspace/nest-grpc-demo/map-service"
	listFiels(dirname, 0)
}

func listFiels(dirname string, level int) {
	// level 用来记录当前递归的层次
	// 生成有层次感的空格
	s := "|--"
	for i := 0; i < level; i++ {
		s = "|  " + s
	}

	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range fileInfos {
		filename := dirname + "/" + fi.Name()
		fmt.Printf("%s%s\n", s, filename)
		if fi.IsDir() {
			// 继续遍历 fi 这个目录
			listFiels(filename, level+1)
		}
	}
}
