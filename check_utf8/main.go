package main

import (
	"bytes"
	"errors"
	"io"
	"log"
	"os"
	"unicode/utf8"
)

var errNonExistentFile = errors.New("file doesn't exist")

// isEncodedInUtf8 io.Readerとして受け取ったファイルの文字コードがutf-8かどうか判定する
func isEncodedInUtf8(file io.Reader) (isUtf8 bool, err error) {
	// io.Readerを[]byteに変換
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, file)
	b := buf.Bytes()
	if len(b) == 0 {
		// ファイルが存在しない場合
		return false, errNonExistentFile
	}
	// 文字コードがutf-8かどうか判定
	return utf8.Valid(b), nil
}

func main() {
	file, err := os.Open("./testdata/utf8_sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	// io.TeeReaderを用いることで同じストリームを使い回せるようにする
	bufFile := bytes.NewBuffer(nil)
	teeFile := io.TeeReader(file, bufFile)

	// ファイルの文字コードがUTF-8どうかを判定
	isUtf8, err := isEncodedInUtf8(teeFile)
	if err != nil {
		log.Fatal(err)
	}
	if !isUtf8 {
		log.Fatal(errors.New("invalid character encoding"))
	}

	buf := bytes.NewBuffer(nil)
	io.Copy(buf, bufFile)
	b := buf.Bytes()
	if len(b) == 0 {
		log.Fatal(errNonExistentFile)
	}

	log.Println("file content: ", string(b))

	log.Println("Done")
}
