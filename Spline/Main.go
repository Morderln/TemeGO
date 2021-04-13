package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func computeDistance(firstPoint Point,secondPoint Point) float64 {
	dist := math.Sqrt(math.Pow(firstPoint.x-secondPoint.x, 2) + math.Pow(firstPoint.y-secondPoint.y, 2) + math.Pow(firstPoint.z-secondPoint.z, 2))
	return dist
}

func computeArea(a, b, c float64) float64 {
	p := (a + b + c) / 2
	return math.Sqrt(p * (p - a) * (p - b) * (p - c))
}

type Point struct {
	x,y,z float64
}


func main() {

	var nrSlices int
	p := 20
	var a,b,c float64
	area:=0.0
	var verteces [] Point

	file, err := os.Open("input.txt")
	if err != nil {
		panic("error")
	}
	output_file, _ := os.Create("result.obj")
	w := bufio.NewWriter(output_file)
	fmt.Fscan(file, &nrSlices)
	for j := 0; j < nrSlices; j++ {

		var x1 []float64
		var y1 []float64

		var x2 []float64
		var y2 []float64
		var nrPuncteDeasupra int
		var nrPuncteDedesubt int

		fmt.Fscan(file, &nrPuncteDeasupra)
		fmt.Fscan(file, &nrPuncteDedesubt)

		var currentX float64
		var currentY float64
		for i := 0; i < nrPuncteDeasupra; i++ {
			fmt.Fscan(file, &currentX)
			fmt.Fscan(file, &currentY)
			x1 = append(x1, currentX)
			y1 = append(y1, currentY)
		}

		for i := 0; i < nrPuncteDedesubt; i++ {
			fmt.Fscan(file, &currentX)
			fmt.Fscan(file, &currentY)
			x2 = append(x2, currentX)
			y2 = append(y2, currentY)
		}

		var s1 = newSpline(x1, y1, CubicSecondDeriv, 0, 0) // val derivatelor de ordin 2 in capete
		// s - setul de functii spline
		var s2 = newSpline(x2, y2, CubicSecondDeriv, 0, 0)

		var p = 20
		var h = (x1[nrPuncteDeasupra-1] - x1[0]) / float64(p)




		for i := 1; i < p; i++ {
			xTemp := x1[0] + float64(i)*h //p pct echidistasnte cu dist=h
			yTemp := s1.At(xTemp)

			fmt.Fprintf(w, "v %f %f %d \n", xTemp, yTemp, j*10)
			verteces=append(verteces,Point{xTemp,yTemp,float64(j*10)})

		}

		for i := 1; i < p; i++ {
			xTemp := x2[0] + float64(i)*h //p pct echidistasnte cu dist=h
			yTemp := s2.At(xTemp)

			fmt.Fprintf(w, "v %f %f %d \n", xTemp, yTemp, j*10)

			verteces=append(verteces,Point{xTemp,yTemp,float64(j*10)})

		}
	}


	p--
	for i := 1; i < p; i++ {
		fmt.Fprintf(w, "f %d %d %d \n", i, 2*p+i, 2*p+i+1)
		a:=computeDistance(verteces[i-1],verteces[2*p+1-1])
		b:=computeDistance(verteces[2*p+1-1],verteces[2*p+1])
		c:=computeDistance(verteces[2*p+1],verteces[i-1])
		area+=computeArea(a,b,c)

		fmt.Fprintf(w, "f %d %d %d \n", i, i+1, 2*p+i+1)
		a=computeDistance(verteces[i-1],verteces[i])
		b=computeDistance(verteces[i],verteces[2*p+1])
		area+=computeArea(a,b,c)

		fmt.Fprintf(w, "f %d %d %d \n", p+i, 3*p+i, 3*p+i+1)
		a=computeDistance(verteces[p+i-1],verteces[3*p+1-1])
		b=computeDistance(verteces[3*p+i-1],verteces[3*p+1])
		c=computeDistance(verteces[3*p+i],verteces[p+i-1])
		area+=computeArea(a,b,c)

		fmt.Fprintf(w, "f %d %d %d \n", p+i, p+i+1, 3*p+i+1)
		a=computeDistance(verteces[p+i-1],verteces[p+1])
		b=computeDistance(verteces[p+i],verteces[3*p+1])
		area+=computeArea(a,b,c)

	}
	p++
	fmt.Fprintf(w, "f %d %d %d \n", 1, p, 3*p-2)
	a=computeDistance(verteces[0],verteces[p-1])
	b=computeDistance(verteces[p-1],verteces[3*p-3])
	c=computeDistance(verteces[3*p-3],verteces[0])
	area+=computeArea(a,b,c)

	fmt.Fprintf(w, "f %d %d %d \n", 1, 2*p-1, 3*p-2)
	a=computeDistance(verteces[0],verteces[2*p-2])
	b=computeDistance(verteces[2*p-2],verteces[3*p-3])
	area+=computeArea(a,b,c)

	fmt.Fprintf(w, "f %d %d %d \n", p-1, 3*p-3, 4*p-4)
	a=computeDistance(verteces[p-2],verteces[3*p-4])
	b=computeDistance(verteces[3*p-4],verteces[4*p-5])
	c=computeDistance(verteces[4*p-5],verteces[p-2])
	area+=computeArea(a,b,c)

	fmt.Fprintf(w, "f %d %d %d \n", p-1, 2*p-2, 4*p-4)
	a=computeDistance(verteces[p-2],verteces[2*p-3])
	b=computeDistance(verteces[2*p-3],verteces[4*p-5])
	area+=computeArea(a,b,c)

	fmt.Println("Area: ",area)

	w.Flush()
}
