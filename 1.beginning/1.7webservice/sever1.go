package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand/v2"
	"net/http"
	"os"
	"strconv"
	"sync"
)

// 互斥锁变量
var mu sync.Mutex
var count int

// 启动后使用1.5节的httpget请求访问测试
// 1.6节的多线程用于count值的更新测试
func main() {
	http.HandleFunc("/", print)
	http.HandleFunc("/count", printCount)
	http.HandleFunc("/printRequest", printRequest)
	http.HandleFunc("/lissajous", lissajousOutput)
	//http.ListenAndServeTLS()，需要经过tls协议认证，需要https访问
	http.ListenAndServe("localhost:8080", nil)
}

// http://localhost:8080/*
func print(w http.ResponseWriter, r *http.Request) {
	//保证修改count值的goroutine同时刻只能有一个
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "%v", r.URL.Path)
	fmt.Printf("received: %s", r.URL.Path)
}

// http://localhost:8080/count
func printCount(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "count: %d\n", count)
}

func printRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

// http://localhost:8080/lissajous?cycle=*
func lissajousOutput(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	fmt.Println("cycle: " + r.Form.Get("cycle"))
	cycle, err := strconv.Atoi(r.Form.Get("cycle"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "err:%v\n", err)
		os.Exit(1)
	}
	lissajous(w, cycle)
}

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex        = 0 // first color in palette
	blackIndex        = 1 // next color in palette
	test_const_string = "123"
)

func lissajous(out io.Writer, cycle int) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	cycles := 5 // number of complete x oscillator revolutions
	if cycle != 0 {
		cycles = cycle
	}

	freq := rand.Float64() * 3.0
	//初始化结构体的方式：
	//1、类似构造函数，Struct{arg1,arg2,...}其中arg*的顺序与结构体定义时一致；
	//2、通过字段名初始化，Struct{Arg1: arg1, Arg2: arg2,...}
	//3、struct := new(Struct) 创建变量后，手动通过struct.arg1 = *赋值
	//note 结构体内变量首字母大写代表公开字段public，小写代表私有字段private，方法可见性同理
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
