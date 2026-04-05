package main

import "fmt"

type User struct {
	Name     string
	Passowrd string
	Age      int
}

func main() {
	hoge := [3]bool{true, true, false}
	// 上記と同様の配列を作成し、それを参照するスライスを作成
	fuga := []bool{true, true, false}
	fuga[2] = true
	fmt.Println(hoge, fuga)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	// structの勉強
	var piyo *User
	piyo = &User{
		Name:     "Taro",
		Passowrd: "password",
		Age:      33,
	}
	fmt.Println(piyo)
	fmt.Println(*piyo)
}
