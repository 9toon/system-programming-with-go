package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)
	if bytes.Equal(buffer, []byte("tEXt")) {
		rawText := make([]byte, length)
		chunk.Read(rawText)
		fmt.Println(string(rawText))
	}
}

func textChunk(text string) io.Reader {
	byteData := []byte(text)
	var buffer bytes.Buffer

	// 「長さ」を書き込み
	binary.Write(&buffer, binary.BigEndian, int32(len(byteData)))
	// 「種類」を書き込み
	buffer.WriteString("tEXt")
	// データ本体を書き込み
	buffer.Write(byteData)
	// CRCを計算
	crc := crc32.NewIEEE()
	io.WriteString(crc, "tEXt")
	// 「CRC」を書き込み
	binary.Write(&buffer, binary.BigEndian, crc.Sum32())
	return &buffer
}

func readChunks(file *os.File) []io.Reader {
	// チャンクを格納する配列
	var chunks []io.Reader

	// 最初の8バイトを飛ばす
	file.Seek(8, 0)
	var offset int64 = 8

	for {
		var length int32
		// 「長さ」を格納している領域を読んでデータ長を取得
		// 現在位置は4バイト先に移動する
		err := binary.Read(file, binary.BigEndian, &length)
		if err == io.EOF {
			break
		}

		// `+12` は、各チャンクで「長さ」「種類」「CRC」を格納するのに各4バイトずつ使うから
		chunks = append(chunks, io.NewSectionReader(file, offset, int64(length)+12))

		// 次のチャンクの先頭に移動
		// 現在位置は「長さ」を読み終わった箇所
		//   -> 種類（4バイト） + データ長 + CRC（4バイト）先に移動する
		offset, _ = file.Seek(int64(length+8), 1)
	}

	return chunks
}

func main() {
	file, err := os.Open("Lenna.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	newFile, err := os.Create("Lenna2.png")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	chunks := readChunks(file)
	// シグネチャを書き込み
	io.WriteString(newFile, "\x89PNG\r\n\x1a\n")

	// 先頭に必要なIHDRチャンクを書き込み
	io.Copy(newFile, chunks[0])

	// テキストチャンクを追加
	io.Copy(newFile, textChunk("ASCII PROGRAMMING++"))

	// 残りのチャンクを書き込み
	for _, chunk := range chunks[1:] {
		io.Copy(newFile, chunk)
	}

	newChunks := readChunks(newFile)
	for _, chunk := range newChunks {
		dumpChunk(chunk)
	}
}
