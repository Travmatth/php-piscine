package main

import (
	Color "d06/ex00"
	Vertex "d06/ex01"
	"fmt"
)

func main() {
	Color.Verbose = false
	doc, err := Vertex.Doc()
	if err == nil {
		fmt.Print(doc)
	}
	Vertex.Verbose = true
	vtx0 := Vertex.NewVertex([]float64{0.0, 0.0, 0.0}, nil)
	defer vtx0.Release()
	fmt.Printf("%+v\n", vtx0)

	red := Color.NewColor([]int{255, 0, 0})
	green := Color.NewColor([]int{0, 255, 0})
	blue := Color.NewColor([]int{0, 0, 255})

	unitX := Vertex.NewVertex([]float64{1.0, 0.0, 0.0}, green)
	unitY := Vertex.NewVertex([]float64{0.0, 1.0, 0.0}, red)
	unitZ := Vertex.NewVertex([]float64{0.0, 0.0, 1.0}, blue)
	defer unitX.Release()
	defer unitY.Release()
	defer unitZ.Release()

	fmt.Printf("%+v\n", unitX)
	fmt.Printf("%+v\n", unitY)
	fmt.Printf("%+v\n", unitZ)

	Vertex.Verbose = false

	sqrA := Vertex.NewVertex([]float64{0.0, 0.0, 0.0}, nil)
	sqrB := Vertex.NewVertex([]float64{4.2, 0.0, 0.0}, nil)
	sqrC := Vertex.NewVertex([]float64{4.2, 4.2, 0.0}, nil)
	sqrD := Vertex.NewVertex([]float64{0.0, 4.2, 0.0}, nil)
	defer sqrA.Release()
	defer sqrB.Release()
	defer sqrC.Release()
	defer sqrD.Release()

	fmt.Println(sqrA)
	fmt.Println(sqrB)
	fmt.Println(sqrC)
	fmt.Println(sqrD)

	A := Vertex.NewVertex([]float64{3.0, 3.0, 3.0}, nil)
	defer A.Release()
	fmt.Println(A)
	equA := Vertex.NewVertex([]float64{9.0, 9.0, 9.0, 3.0}, nil)
	defer equA.Release()
	fmt.Println(equA)
	Vertex.Verbose = true
}
