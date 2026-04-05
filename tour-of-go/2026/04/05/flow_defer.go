package main

import "fmt"

func main() {
	// deferステートメントはdeferへ渡した関数の実行を
	// 呼び出し元の関数の終わり（returnする）まで遅延させる
	defer fmt.Println("world")

	fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
