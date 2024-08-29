package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().UnixMilli()
	for i := 0; i < 1000000; i++ {
		PopCount1(uint64(i))
	}
	fmt.Printf("PopCount1 :%d \n", time.Now().UnixMilli()-t)
	t = time.Now().UnixMilli()
	for i := 0; i < 1000000; i++ {
		PopCount2(uint64(i))
	}
	fmt.Printf("PopCount2 :%d \n", time.Now().UnixMilli()-t)
}

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount1(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount2(x uint64) int {
	var result int
	for x > 0 {
		result += int(x & 1)
		x >>= 1
	}
	return result
}
