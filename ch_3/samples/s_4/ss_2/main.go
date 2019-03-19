package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	// defer: 現在のスコープを抜けたときに、後ろに書いた処理が必ず実行される
	defer file.Close()
	io.Copy(os.Stdout, file)
}
