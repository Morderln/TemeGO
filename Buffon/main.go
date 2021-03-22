package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	var n float64
	var L float64
	L = 1
	var nr float64
	nr = 0
	n = 100000
	for i:=1; i<=int(n);i++{
		y:=rand.Float64()*float64(100*L)
		alfa:=rand.Float64()*math.Pi/2
		var y1 = y + float64(L)*math.Sin(alfa)
		b := int(y/L) != int(y1/L)
		if b {
			nr++
		}
	}


	fmt.Println(2.*n/nr)
}