package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(add(5, 5))

	var i, j int = 1, 2
	k := 3
	var c, python bool = false, false
	fmt.Println(i, j, k, c, python)

	defer fmt.Print(", suzuki\n")

	const (
		PI  = 3.14
		PI2 = 3.1415
	)
	fmt.Println(PI + float64(i))

	defer fmt.Printf("world")

	if v := math.Pow(2, 10); v < 10000 {
		fmt.Println(v)
	}
	if false {
	} else {
		fmt.Print("hello, ")
	}

	foo := 11111
	p := &foo
	fmt.Println(*p)
}

func add(x, y int) int {
	return x + y
}
