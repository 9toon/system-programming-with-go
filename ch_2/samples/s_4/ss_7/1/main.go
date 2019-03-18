package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// %v は何でも表示できるフォーマット指定子で
	// どんな型でも、String()メソッドがあればそれを使って出力する
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v", time.Now())
}
