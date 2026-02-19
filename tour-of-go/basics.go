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

// Variables
// var c, python, java bool

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(math.Pi);
	fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))

	var i, j int = 1, 2
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
	fmt.Printf("c: %T, python: %T, java: %T\n", c, python, java)

	// 関数の中では暗黙的な型宣言が使える
	k := 3 // var 宣言の代わり
	fmt.Println(i, j, k)
}	

