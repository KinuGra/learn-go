package main

import "fmt"

const Pi float64 = 3.14

// 型のない定数を使う際に必要な型として扱える
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func main() {
	n1 := 100 // 右側の変数から型推論
	n2 := 3.1415
	n3 := 50 + 20i
	fmt.Printf("%T, %T, %T\n", n1, n2, n3)

	// 定数（文字、文字列、boolean、数値のみ）
	const Truth = true
	const World = "世界"
	fmt.Println("Hello, ", World)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}
