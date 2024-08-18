package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	//创建一个名为ch的通道，传输string值
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		//启动一个并发执行单元 goroutine
		go fetch(url, ch)
	}
	//这里循环 args 次是因为，fetch函数确保了，每次执行fetch只有一次输出到ch（无论是否报错），
	//如果一次fetch中存在多次ch的写入操作，则会导致后面的httpget返回无法通过ch输出；
	//抑或者，一次fetch中不一定有一次输出，则会导致程序阻塞在该点。
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	//ioutil.Discard是一个特殊的io.Writer接口，丢弃所有写入数据（这里是供测试，只考虑字节数）
	data, err := io.ReadAll(resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		//Sprintf是构造字符串（并非直接打印出来）
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	//消耗了多少时间
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("time consumed: [%.2fs]  data: [%s]", secs, data)
}
