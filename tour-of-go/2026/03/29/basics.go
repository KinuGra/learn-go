package main

import (
	"fmt"
	f "fmt"
	"math"
	"time"
)

var (
	c int
	b int
)

/*
func ex_func
{
	この書き方はできない
}
*/

// 変数に初期値を与えずに宣言するとゼロ値が与えられます
// string: "", bool: false, int,float: 0

func main() {
	fmt.Println("hello, world")
	fmt.Println("This time is ", time.Now())
	fmt.Println(math.Pi)
	a, b, c := a()
	f.Println(a, b, c)
	x := "hello"
	y := "world"
	x, y = swap(x, y)
	f.Println(x, y)
	d, e := hoge()
	f.Println(d, e)

	var python, cs, cpp bool
	var i int
	f.Println(python, cs, cpp, i)

	var j, k int = 2, 3
	l := 10.5
	f.Print(j, k, l, "\n")
}

func a() (float32, float32, float32) {
	return 1.2, 3.52, 8.221
}

func swap(x, y string) (string, string) {
	return y, x
}

func hoge() (x int, y string) {
	x = 10
	y = "hogehoge"
	return
}

