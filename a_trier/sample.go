package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

// Sample1 selects a single line with uniform probability from a stream.
func Sample1(in *bufio.Reader) []byte {
	var (
		r []byte
		s = int64(1)
	)
	for {
		b, err := in.ReadBytes('\n')
		if b == nil || err == io.EOF {
			break
		}
		if rand.Int63()%s == 0 {
			// select 'b' as 'r' with probability 1/s.
			// Instead of sampling in [0,1] and comparing to 1/s,
			// sample an int modulo 's' and compare to 0.
			r = b
		}
		s++
	}
	return r
}

// SampleN selects N lines with uniform probability from a stream.
func SampleN(n int, in *bufio.Reader) [][]byte {
	var (
		r = make([][]byte, n)
		s = int64(1)
		m = int64(n)
	)
	// fill reservoir
	for k := range r {
		b, err := in.ReadBytes('\n')
		if b == nil || err == io.EOF {
			return r[:k] // early return
		}
		r[k] = b
		s++
	}
	// sample steam
	for {
		b, err := in.ReadBytes('\n')
		if b == nil || err == io.EOF {
			break
		}
		if k := rand.Int63() % s; k < m {
			r[k] = b
		}
		s++
	}
	return r
}

func SampleN2(n int, in *bufio.Reader) [][]byte {
	itor := &BytesIter{r: *in}
	itee := InitSamplerN(n)
	Consume(itor, itee)
	itee.Finish()
	return itee.Get()
}

func Sample2(in *bufio.Reader) []byte {
	itor := &BytesIter{r: *in}
	itee := &Sampler{}
	Consume(itor, itee)
	return itee.r
}

func main() {
	rand.Seed(time.Now().UnixNano())

	r := bufio.NewReader(os.Stdin)

	//fmt.Println(string(Sample1(r)))

	for _, b := range SampleN2(10, r) {
		s := string(b[:len(b)-1])
		fmt.Println(s)
	}
}

type Iteratee interface {
	Push(b []byte) bool // ignore nil
	Finish() error
}

type Iterator interface {
	Next() bool
	Get() []byte
}

func Consume(itor Iterator, itee Iteratee) {
	for itor.Next() {
		itee.Push(itor.Get())
	}
}

type BytesIter struct {
	r bufio.Reader // try embedding
	b []byte
}

// Next returns a line from the given buffered Reader, or returns nil when EOF
// is reached, or panic on errors.
func (iter *BytesIter) Next() bool {
	b, err := iter.r.ReadBytes('\n')
	switch {
	case err == nil:
		iter.b = b
		return true
	case err == io.EOF:
		iter.b = nil
		return false
	default:
		panic(err)
	}
}

func (iter *BytesIter) Get() []byte {
	return iter.b
}

// SamplerN selects N lines with uniform probability from a stream.
type SamplerN struct {
	r [][]byte
	n int64
	k int64
}

func InitSamplerN(n int) *SamplerN {
	return &SamplerN{
		r: make([][]byte, n),
		n: int64(n),
		k: 0,
	}
}

func (s *SamplerN) Push(b []byte) bool {
	s.k++
	switch {
	case b == nil:
		s.k--
	case s.k < s.n:
		s.r[s.k] = b
	default:
		if j := rand.Int63() % s.k; j < s.n {
			s.r[j] = b
		}
	}
	return true
}

func (s *SamplerN) Finish() error {
	s.r = s.r[:min(s.k, s.n)]
	s.k, s.n = 0, 0
	return nil
}

func (s *SamplerN) Get() [][]byte {
	return s.r
}

type Sampler struct {
	r []byte
	k int64
}

func (s *Sampler) Push(b []byte) bool {
	switch {
	case b == nil:
		return true
	case rand.Int63()%s.k == 0:
		s.r = b
	}
	s.k++
	return true
}

func (s *Sampler) Finish() error { return nil }

func min(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
