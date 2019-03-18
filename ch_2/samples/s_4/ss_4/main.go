// 書かれた内容を記憶しておくバッファ（2）：  strings.Builder
// バッファ経由で出力（Go v1.10で新しく入った機能使った版）
package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder
	builder.Write([]byte("strings.Builder example\n"))
	fmt.Println(builder.String())
}
