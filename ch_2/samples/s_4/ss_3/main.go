// 書かれた内容を記憶しておくバッファ（1）： bytes.Buffer
// バッファ経由で標準出力に出力
package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buffer.String())
}
