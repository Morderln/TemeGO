package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Introduceti raza: ")
	var r float64
	fmt.Scan(&r)

	fmt.Println("Introduceti nr de trageri:")
	var n float64
	fmt.Scan(&n)

	pi:=0.0
	min:=0.0
	max:=2.0*r
	m:=0.0
	for i:=0; i<int(n);i++ {
		x:=(rand.Float64()* (max-min))+min
		y:=(rand.Float64()* (max-min))+min
		if math.Pow((x-r),2) + math.Pow((y-r),2) <=math.Pow(r,2) {
			m+=1
		}
	}
	fmt.Println("Trageri:",m)
	pi = 4.0 * m / n
	fmt.Printf("%.6f",pi)
}
