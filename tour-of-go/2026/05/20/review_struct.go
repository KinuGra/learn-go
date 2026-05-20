package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	v1 := Vertex{1, 2}
	v2 := Vertex{Y: 1}
	v3 := Vertex{}
	v4 := &Vertex{3, 4}

	fmt.Println(v1, v2, v3, v4, *v4, (*v4).X, v4.X) // v4.X: 自動でデリファレンス
}
