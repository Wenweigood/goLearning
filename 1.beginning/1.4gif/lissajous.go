package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	//控制台输出
	// lissajous(os.Stdout)
	file, _ := os.Create("lissajous.gif")
	//os.file实现了io.Writer接口（隐式实现，仅需要有符合该接口定义的方法即可）
	lissajous(file)
	file.Close()
}

// 定义切片（类似数组，长度：当前元素数，容量：最大元素数）
var palette = []color.Color{color.White, color.Black}

// 定义常量
const (
	whiteIndex        = 0 // first color in palette
	blackIndex        = 1 // next color in palette
	test_const_string = "123"
)

// go语言变量命名基本遵循{变量名 变量类型}的格式，如 var a int = 123，
// 与c++、java基本相反，包括指针命名c++:{int*}、go:{*int}
func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

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
		for t := 0.0; t < cycles*2*math.Pi; t += res {
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
