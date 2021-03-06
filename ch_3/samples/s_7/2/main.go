package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	var buffer bytes.Buffer
	reader := bytes.NewBufferString("Example of io.TeeReeder\n")
	teeReader := io.TeeReader(reader, &buffer)

	// データを読み捨てる
	_, _ = ioutil.ReadAll(teeReader)

	// が、バッファには残っている
	fmt.Println(buffer.String())
}
