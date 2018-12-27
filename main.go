package main

import (
	"fmt"
	"github.com/totoleo/tero/parse"
	"os"
)

/*
Kitc client 生成改进方案
必要的输入：
1. Idl 文件目录，默认位置：$HOME/idl/
2. 被调用方的 psm
3. 调用方的 psm , 默认是 clients/../scripts/setting.py 内的配置
4. 服务的别名
5. 生成文件的名称，默认是: {psm}_client.go {psm}_client_test.go
6. 输出目录：clients/
7.
*/
func main() {
	file := os.Getenv("GOFILE")

	fileSet, err := parse.File(file, true)
	fmt.Println(fileSet, err)
	if err != nil {
		panic(err)
	}
	for k, set := range fileSet.Specs {
		fmt.Println(k, set.Pos(), set.End())
	}

}
