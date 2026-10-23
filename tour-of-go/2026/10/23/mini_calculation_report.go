package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/*
復習問題：ミニ計算レポート
次の仕様を満たす実行可能なGoプログラムを書いてください。

- あいさつを表示する処理を、独立した関数として作り、プログラムから呼び出す。
- 現在時刻と、0〜9のランダムな整数を表示する。
- 8の平方根と円周率を、簡潔な浮動小数点表記で表示する。
- 1行で、文字列・整数の型・整数の2進表記・小数点以下2桁の小数を表示する。
- 2つの整数の合計を返す関数と、3つの整数の合計を返す関数を作り、それぞれ 22, 28 と 3, 100, 7 で呼び出して結果を表示する。
- 2つの文字列を受け取り、順序を入れ替えた2つの値を返す関数を作る。"hello" と "world" を渡して、返された値を表示する。
- 必要なパッケージは自分で選び、関数名・引数・戻り値も自分で設計する。
*/

func main() {
	greet()
	fmt.Println("現在時刻:", time.Now())
	fmt.Println(rand.Intn(10))
	fmt.Printf("8の平方根: %g, 円周率: %g\n", math.Sqrt(8), math.Pi)
	fmt.Printf("%s、%T、%b、%.2f\n", "hello", 100, 31, 3.141592)
	fmt.Println(add(22, 28))
	fmt.Println(calc(3, 100, 7))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

func greet() {
	fmt.Println("こんにちは")
}

func add(x int, y int) int {
	return x + y
}

func calc(x, y, z int) int {
	return x + y + z
}

func swap(s1, s2 string) (string, string) {
	return s2, s1
}
