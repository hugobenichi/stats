package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"io"
	"os"
)

// TODO:
//	- add numbers
//  - add reverse print
//  - add topor bottom option
//  - add option to change k

func main() {
	k := 10
	topk := topTokens(os.Stdin, k)
	for i := 0; i < k; i++ {
		fmt.Println(string(heap.Pop(&topk).([]byte)))
	}
}

func topTokens(source io.Reader, k int) TokSlice {
	var (
		r             = bufio.NewReader(source)
		topk TokSlice = make([][]byte, 0, k+1)
	)
	for {
		bytes, err := r.ReadBytes('\n')
		if err != nil { // EOF
			return topk
		}
		heap.Push(&topk, bytes[:len(bytes)-1])
		if len(topk) > k {
			heap.Pop(&topk)
		}
	}
}

type TokSlice [][]byte

func (t *TokSlice) Len() int           { return len(*t) }
func (t *TokSlice) Less(i, j int) bool { return bytes.Compare((*t)[i], (*t)[j]) < 1 }
func (t *TokSlice) Swap(i, j int)      { (*t)[i], (*t)[j] = (*t)[j], (*t)[i] }
func (t *TokSlice) Push(x interface{}) { *t = append(*t, x.([]byte)) }
func (t *TokSlice) Pop() interface{} {
	l := len(*t) - 1
	e := (*t)[l]
	*t = (*t)[:l]
	return e
}
