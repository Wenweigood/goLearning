package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//建立一个从string->int的映射
	counts := make(map[string]int)
	//标准输入std:in
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		//map会自动处理已存在的映射，和未存在的映射关系，未存在的映射返回的是value的默认值（int，0）
		counts[input.Text()]++
	}
	//在counts这个map上迭代，line 是key，n是value
	for line, n := range counts {
		fmt.Println(line)
		fmt.Println(n)
	}
}

/* printf中的占位符
%d          十进制整数
%x, %o, %b  十六进制，八进制，二进制整数。
%f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
%t          布尔：true或false
%c          字符（rune） (Unicode码点)
%s          字符串
%q          带双引号的字符串"abc"或带单引号的字符'c'
%v          变量的自然形式（natural format）
%T          变量的类型
%%          字面上的百分号标志（无操作数）
*/
