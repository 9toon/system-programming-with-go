package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	// ヘッダーを表示
	fmt.Println(res.Header)

	// 最後に閉じるの忘れない
	defer res.Body.Close()

	// ボディを表示
	io.Copy(os.Stdout, res.Body)
}
