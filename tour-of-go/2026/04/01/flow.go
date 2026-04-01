package main

import (
	"fmt"
	"math"
)

func main() {
	sum := 0
	// ()はないが、{}は常に必要
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			fmt.Printf("%v ", i*j)
		}
		fmt.Println()
	}

	// 初期化と後処理ステートメントの記述は任意です
	x := 1
	for x < 1000 {
		x += 1
		fmt.Println(x)
	}

	// Goではwhileはないのでforだけを使う
	for x > 0 {
		x--
		fmt.Println(x)
	}

	for {
		x++
		fmt.Println(x)
		if x == 10000 {
			break
		}
	}

	// ifステートメントはfor文と同様に()不要で{}は必要
	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
