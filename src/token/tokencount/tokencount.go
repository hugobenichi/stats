package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unsafe"
)

func main() {
	for tok, count := range tokens(os.Stdin) {
		tok := tok[:len(tok)-1]
		fmt.Println(tok, count)
	}
}

func tokens(source io.Reader) map[string]int {
	r := bufio.NewReader(source)
	toks := make(map[string]int)
	for {
		bytes, err := r.ReadBytes('\n')
		if err != nil { // EOF
			return toks
		}
		line := unsafeString(bytes)
		toks[line]++
	}
}

func unsafeString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
