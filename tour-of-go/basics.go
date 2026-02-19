package main

import (
	"fmt"
	"math/rand"
	"math"
)

// プログラムはmainパッケージから開始される
// 最初の文字が大文字で始まる花江は外部のパッケージから参照できるエクスポートされた名前
// Piはmathパッケージでエクスポート（公開）されている

func add(x int, y int) int { // x, y intに省略できる
	return x + y
}

// 複数の戻り値を返せる
func swap(x, y string) (string, string) {
	return y, x
}

// 戻り値に名前をつける
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(math.Pi);
	fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
}

