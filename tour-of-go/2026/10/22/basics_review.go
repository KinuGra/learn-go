package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	hello()
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(8))
	fmt.Printf("%g\n", math.Pi)
	fmt.Printf("%v, 型: %T, %b, %.2f\n", "汎用", 100, 1200, 3.141592)
	fmt.Println(add(22, 28))
	fmt.Println(calc(3, 100, 7))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

func hello() { fmt.Println("hello, world") }
func add(x int, y int) int {
	return x + y
}
func calc(x, y, z int) int { return x + y + z } // 2つ以上の引数が同じ型であれば最後の型を残して省略可
func swap(a, b string) (string, string) {
	return b, a // 関数は複数の戻り値を返せる
}

// basics/6まで
