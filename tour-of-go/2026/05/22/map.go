package main

import "fmt"

// 緯度、経度
type Vertex struct {
	Lat, Long float64
}

var (
	m      map[string]Vertex
	person map[string]int
)

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	person = make(map[string]int)
	person["suzuki"] = 150
	person["yokota"] = 200
	for k, v := range person {
		fmt.Println(k, v)
	}

	// mapリテラル
	ml := map[string]Vertex{
		"Bell": { // "Bell": Vertex{}のVertexは省略できる
			3.14, 100.50,
		},
		"Google": {
			37.4, -122.08,
		},
	}
	fmt.Println(ml)
}
