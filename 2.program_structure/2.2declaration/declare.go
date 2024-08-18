package main

import (
	"fmt"
	"strconv"
)

// 常量声明
const base int = 1

func main() {
	//变量声明
	var sale int = 2

	fmt.Println(sale * base)
	fmt.Println(int2string(sale))
}

// 函数声明
func int2string(i int) string {
	return strconv.Itoa(i)
}

// 类型声明，声明现有类型，起别名作用
type myInt int

// 类型声明，声明新复杂类型，如结构体struct、接口interface、切片、映射
type MyType struct {
	Name string
	Age  int
}
