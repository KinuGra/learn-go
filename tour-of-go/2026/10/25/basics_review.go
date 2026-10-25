package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// varステートメントは変数を宣言
var c, python, kotlin bool

// x := 100 => 関数の外ではvar, funcのようなキーワードで始まる宣言が必要

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Println(split(17))

	var i int
	fmt.Println(i, c, python, kotlin)

	// 初期化子
	var j, k float64 = 3.14, 1.414
	cpp, rust, r := true, "Yes!", 10 // 暗黙的な型宣言
	fmt.Println(j, k, cpp, rust, r)
	l, m := 3, "hello" // 短縮変数宣言、初期化必須
	fmt.Printf("%T, %T\n", l, m)

	// テスト
	fmt.Printf("%v\n", hoge())

	// Goの基本型（組み込み型）
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// 初期値を与えなければゼロ値が与えられる
	var x int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", x, f, b, s)

	// Cとは違い、Goでの型変換は明示的な変換が必要
	var y int = 100
	var f2 float64 = math.Sqrt(float64(y))
	var u uint = uint(f2) // var u uint = f2ではエラー
	fmt.Println(y, f2, u)
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // naked return
}

func hoge() (hoge bool) {
	c = true
	hoge = c
	return
}
