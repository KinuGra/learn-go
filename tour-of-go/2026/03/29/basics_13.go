package main

import "fmt"

// 型変換
func main() {
	var a uint = 15
	f := float64(a)
	fmt.Println(f)
	f = 3.14
	fmt.Println(int(f))
	f = 3
	a = uint(f)
}
