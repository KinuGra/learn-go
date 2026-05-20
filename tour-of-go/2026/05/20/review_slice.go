package main

import "fmt"

func main() {
	// 配列
	var a [3]int
	b := [3]string{"abc", "dce"}
	c := [...]int{1, 2, 3}
	fmt.Println(len(a), a, len(b), b, len(c), c)

	// スライス
	/* スライス同士は裏配列を共有する */
	s1 := []int{1, 2, 3} // リテラル
	s2 := make([]int, 5)
	s3 := make([]int, 0, 10)
	fmt.Println(cap(s1), cap(s2), cap(s3))
	fmt.Println(s1, s2, s3)

	arr := [5]int{10, 20, 30, 40, 50}
	s4 := arr[1:4]
	fmt.Println(s4)

	s := []int{1, 2}
	s = append(s, 3, 4)
	fmt.Println(s)
	s5 := make([]int, len(s))
	copy(s5, s)
	fmt.Println(s5)

	var s6 []int
	s6 = append(s6, 1, 2, 3, 4, 5)
	fmt.Println(s6)
}
