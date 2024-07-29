package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var result string
	for i := 1; i < len(os.Args); i++ {
		result += os.Args[i] + ","
	}
	result = strings.TrimSuffix(result, ",")
	fmt.Printf("args is : %s \n", result)

	fmt.Println("hello, world!")
	fmt.Printf("text:%s, hour:%d, minute:%d, second:%d", time.Now().String(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())
}
