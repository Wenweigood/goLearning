package main

import (
	"fmt"
	//导入语句将导入的包绑定到一个短小的名字，然后通过该短小的名字就可以引用包中导出的全部内容
	tempconv "goLearning/2.program_structure/2.6package"
	"os"
	"strconv"
)

func main() {
	fmt.Println(tempconv.C2F(tempconv.BoillingC))

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.F2C(f), c, tempconv.C2F(c))
	}
}
