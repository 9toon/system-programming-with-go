package main

import "fmt"

// インターフェースを定義
type Talker interface {
	Talk()
}

// 構造体を定義
type Greeter struct {
	name string
}

// 構造体はインターフェースで定義されているメソッドを持つ
func (g Greeter) Talk() {
	fmt.Printf("Hello, my name is %s\n", g.name)
}

func main() {
	// インターフェースの型を持つ変数を宣言
	var talker Talker
	// インターフェースを満たす構造体のポインタは代入できる
	talker = &Greeter{"wozozo"}
	talker.Talk()
}
