package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/*
復習問題：変数と型
このファイルに、次の要件を満たす実行可能なプログラムを書いてください。
関数名・引数・戻り値の型や、変数の宣言方法は自分で決めてください。

- 初期値を指定せずに、パッケージレベルで複数の真偽値を宣言し、初期値を表示する。
- そのうち1つを別の関数内で変更し、変更後の値を返して表示する。
- 整数 17 を受け取り、整数計算で「17の4/9」と「残り」を求めて、2つの値を返す関数を作る。
名前付き戻り値と、戻り値を省略する形式を使う。
- 型を明示する宣言、型推論を使う宣言、短縮形式の宣言をそれぞれ使い、値と型を表示する。
- 整数・浮動小数点数・真偽値・文字列のゼロ値を表示する。文字列は引用符付きで表示する。
- 64ビット符号なし整数で表せる最大値を作り、その値と型を表示する。
- 複素数 -5 + 12i の平方根を求め、値と型を表示する。
- 整数 100 を浮動小数点数に変換して平方根を求め、その結果を符号なし整数に変換して、値と型を表示する。
*/

var (
	c, java, python bool
	h               = 3.14
)

func main() {
	fmt.Println(c, java, python)

	fmt.Println(bar())

	fmt.Println(split(17))

	var hoge string
	foo := 5 + 2i
	fuga := 1.414
	fmt.Printf("%s %T, %v %T, %v %T\n", hoge, hoge, foo, foo, fuga, fuga)

	var q int64
	var w float32
	var e string
	fmt.Printf("%d, %g, %q, %v\n", q, w, e, java)

	var n uint64 = 1<<64 - 1
	fmt.Printf("%v, %T\n", n, n)

	// 複素数の平方根
	z := cmplx.Sqrt(-5 + 12i)
	fmt.Printf("値：%v, 型：%T\n", z, z)

	r := 100
	t := math.Sqrt(float64(r))
	var y uint = uint(t)
	fmt.Printf("値：%v, 型：%T\n", y, y)

	var m int32 = 1<<31 - 1
	fmt.Println(m)

	fmt.Printf("%v, %T\n", h, h)
}

func split(a int) (l, r int) {
	l = a * 4 / 9
	r = a - l
	return
}

func bar() bool {
	c = true
	return c
}
