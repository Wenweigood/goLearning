package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello, world")
	fmt.Printf("text:%s, hour:%d, minute:%d, second:%d", time.Now().String(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())
}
