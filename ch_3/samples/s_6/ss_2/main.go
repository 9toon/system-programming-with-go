package main

import (
	"fmt"
	"strings"
)

var source1 = "123 1.234 1.0e4 test"
var source2 = "123, 1.234, 1.0e4, test"

func main() {
	reader1 := strings.NewReader(source1)
	var i int
	var f, g float64
	var s string
	// Fscan() はスペース区切りになっているのが前提
	fmt.Fscan(reader1, &i, &f, &g, &s)
	fmt.Printf("i=%#v f=%#v g=%#v s=%#v\n\n", i, f, g, s)

	reader2 := strings.NewReader(source2)
	// Fscanf() は任意のパターンを指定できる
	// ↓ ではカンマ+スペース区切り
	fmt.Fscanf(reader2, "%v, %v, %v, %v", &i, &f, &g, &s)
	fmt.Printf("i=%#v f=%#v g=%#v s=%#v\n\n", i, f, g, s)
}
