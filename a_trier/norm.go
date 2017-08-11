package main

import (
	"fmt"
	"math"
	"math/rand"
)

func accept(x, y float64) bool {
	return math.Exp(-x*x/2) > y
}

func main() {
	for {
		if x, y := rand.Float64()*10-5, rand.Float64(); accept(x, y) {
			fmt.Println(x)
		}
	}
}
