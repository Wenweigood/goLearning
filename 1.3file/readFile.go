package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		fmt.Println(string(filename))
		//数据以及读取时的报错
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		//index value，其中index不需使用因此使用_忽略（不能以其他变量名代替，因为存在未使用变量会导致编译出错）
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
		for line, n := range counts {
			if n >= 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}
}
