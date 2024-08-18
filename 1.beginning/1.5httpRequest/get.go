package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			//%v代表一个通用的占位符
			fmt.Fprintf(os.Stderr, "err:%v\n", err)
			os.Exit(1)
		}
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "err:%v\n", err)
			os.Exit(1)
		}
		fmt.Print(string(body) + "\n")
	}
}
