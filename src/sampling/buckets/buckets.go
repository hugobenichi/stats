package main

import (
	"errors"
	"fmt"
	"math"
)

// command line interface example:
//
//  specify by parameters:
//		"fist 0.1, last 100, base 10, segments 4, fraction 4"
//
//  by description:
//		"from 0.1 to 100 by 10:4"  => first 0.1, last 100, base 10, segments 4, fraction 4
//
//  "by octave" == "by 8"
//  "by decade" == "by 10"
//
//  "log from 0.1 by 10:4 in 15"
//  "log to 100 by 10:5  15"
//
//
//
//  grammar
//  specs = type * values
//  values = term * values | ""
//  term = "from" * float			// left boundary
//       | "to" * float				// right boundary
//       | "by" * float":"int // base and number of divisions in the base
//       | "in" * float       // number of buckets
//			 | ???  * float       // number of segments
//
// hiearchie of parameters ?

func main() {
	bs := Log{
		Base: 10, //0.1,
		//First:     0.003, //3,
		Last:      30,
		Segments:  4,
		Divisions: 2,
	}.Buckets()
	for _, b := range bs {
		fmt.Println(b)
	}

	cs := Lin{
		Width: 1, //0.1,
		//First:     4, //3,
		Last:      8,
		Segments:  4,
		Divisions: 4,
	}.Buckets()
	for _, b := range cs {
		fmt.Println(b)
	}
}

type Bucketable interface {
	Buckets() []float64
}

type Log struct {
	Base      float64 `name:"base" default:"10" `
	First     float64
	Last      float64
	Segments  uint
	Divisions uint
}

func (l Log) Buckets() []float64 {
	b := make([]float64, l.Segments*l.Divisions+1)
	switch {
	case l.First != 0:
		for i := range b {
			b[i] = l.First * math.Pow(l.Base, float64(i)/float64(l.Divisions))
		}
	case l.Last != 0:
		j := len(b) - 1
		for i := range b {
			b[j-i] = l.Last * math.Pow(l.Base, -float64(i)/float64(l.Divisions))
		}
	default:
		panic(errors.New("no .First or .Last set"))
	}
	return b
}

type Lin struct {
	Width     float64
	First     float64
	Last      float64
	Segments  uint
	Divisions uint
}

func (l Lin) Buckets() []float64 {
	b := make([]float64, l.Segments*l.Divisions+1)
	switch {
	case l.First != 0:
		for i := range b {
			b[i] = l.First + l.Width*float64(i)/float64(l.Divisions)
		}
	case l.Last != 0:
		j := len(b) - 1
		for i := range b {
			b[j-i] = l.Last - l.Width*float64(i)/float64(l.Divisions)
		}
	default:
		panic(errors.New("no .First or .Last set"))
	}
	return b
}
