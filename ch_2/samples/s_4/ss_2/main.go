// 画面出力
package main

import (
	"os"
)

func main() {
	os.Stdout.Write([]byte("os.Stdout example\n"))
}
