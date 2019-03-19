package main

import (
	"bufio"
	"fmt"
	"strings"
)

var source1 = `1行目
2行目
3行目`
var source2 = "This is a pen"

func main() {
	scanner1 := bufio.NewScanner(strings.NewReader(source1))
	for scanner1.Scan() {
		fmt.Printf("%#v\n", scanner1.Text())
	}

	scanner2 := bufio.NewScanner(strings.NewReader(source2))
	// 分割関数を指定することで改行以外にも対応可能
	scanner2.Split(bufio.ScanWords)
	for scanner2.Scan() {
		fmt.Printf("%#v\n", scanner2.Text())
	}
}
