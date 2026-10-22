## Goのpackageとmain

- **package** は、Goのコードをまとめる単位。同じディレクトリのGoファイルは通常、同じpackageに属する。
- 実行可能なプログラムは `package main` とし、実行開始点として `func main()` を定義する。ファイル名が `main.go` である必要はない。
- 自作packageを別packageから使う場合は、通常は別ディレクトリに置く。
- `go.mod` の `module` 名とディレクトリ名をつなげたパスでimportする。
- 他packageから使う関数・型の名前は大文字で始める。
