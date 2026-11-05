package main

import "fmt"

// Structs
type Vertex struct {
	X int
	Y int
}

func main() {
	// C言語と違い、ポインタ演算はない
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	p2 := &v
	p2.X = 1e9
	(*p2).Y = 0
	fmt.Println(v)

	// Struct Literals
	v1 := Vertex{1, 2}
	v2 := Vertex{Y: 2, X: 4}
	v3 := Vertex{}
	var p3 *Vertex = &Vertex{1, 2}
	fmt.Printf("%v: %p, %v: %p\n%v: %p, %v: %p\n", v1, &v1, v2, &v2, v3, &v3, *p3, p3)

	// Arrays（配列の長さは型の一部分ですので変更できません）、固定長
	var a [2]string
	a[0] = "hello,"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Slices（可変長）、配列よりも一般的、[low:high)、配列への参照のようなもの
	var s []int = primes[1:4]
	fmt.Println(s)
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	n1 := names[0:2]
	n2 := names[1:3]
	fmt.Println(n1, n2)

	n2[0] = "XXX"
	fmt.Println(n1, n2)
	fmt.Println(names) // 元となる配列を共有しているので配列が変更される

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	ss := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
	}
	fmt.Println(ss)
}
