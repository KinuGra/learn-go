package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// 長さdyの外側スライスを作成
	picture := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		// 各行に長さdxの内側スライスを割り当て
		picture[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			// ピクセル値を計算
			picture[y][x] = uint8((x + y) / 2)
		}
	}
	return picture
}

func main() {
	pic.Show(Pic)
}
