package main

import (
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"unicode/utf16"

	"golang.org/x/crypto/md4"
)

var (
	data string
)

func init() {
	flag.StringVar(&data, "d", "input string", "")
	flag.Parse()
}

func main() {
	h := md4.New()
	binary.Write(h, binary.LittleEndian, utf16.Encode(([]rune)(data)))
	fmt.Println(hex.EncodeToString(h.Sum(nil)))
}
