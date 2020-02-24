package main

import (
	Vertex "d06/ex01"
	Vector "d06/ex02"
	"fmt"
)

func main() {
	Vertex.Verbose = false

	doc, err := Vector.Doc()
	if err == nil {
		fmt.Println(doc)
	}
	Vector.Verbose = true

	vtxO := Vertex.NewVertex([]float64{0.0, 0.0, 0.0}, nil)
	vtxX := Vertex.NewVertex([]float64{1.0, 0.0, 0.0}, nil)
	vtxY := Vertex.NewVertex([]float64{0.0, 1.0, 0.0}, nil)
	vtxZ := Vertex.NewVertex([]float64{0.0, 0.0, 1.0}, nil)

	vtcXunit := Vector.NewVector(vtxO, vtxX)
	defer vtcXunit.Release()
	vtcYunit := Vector.NewVector(vtxO, vtxY)
	defer vtcYunit.Release()
	vtcZunit := Vector.NewVector(vtxO, vtxZ)
	defer vtcZunit.Release()

	fmt.Printf("%+v\n", vtcXunit)
	fmt.Printf("%+v\n", vtcYunit)
	fmt.Printf("%+v\n", vtcZunit)

	dest1 := Vertex.NewVertex([]float64{-12.34, 23.45, -34.56}, nil)
	Vertex.Verbose = true
	vtc1 := Vector.NewVector(nil, dest1)
	Vertex.Verbose = false

	orig2 := Vertex.NewVertex([]float64{23.87, -37.95, 78.34}, nil)
	dest2 := Vertex.NewVertex([]float64{-12.34, 23.45, -34.56}, nil)
	vtc2 := Vector.NewVector(orig2, dest2)

	fmt.Printf("Magnitude is %+v\n", vtc2.Magnitude())

	nVtc2 := vtc2.Normalize()
	fmt.Printf("Normalized $vtc2 is %+v\n", nVtc2)
	fmt.Printf("Normalized $vtc2 magnitude is %+v\n", nVtc2.Magnitude())

	fmt.Printf("$vtc1 + $vtc2 is %+v\n", vtc1.Add(vtc2))
	fmt.Printf("$vtc1 - $vtc2 is %+v\n", vtc1.Sub(vtc2))
	fmt.Printf("opposite of $vtc1 is %+v\n", vtc1.Opposite())
	fmt.Printf("scalar product of $vtc1 and 42 is %+v\n", vtc1.ScalarProduct(42))
	fmt.Printf("dot product of $vtc1 and $vtc2 is %+v\n", vtc1.DotProduct(vtc2))
	fmt.Printf("cross product of $vtc1 and $vtc2 is %+v\n", vtc1.CrossProduct(vtc2))
	fmt.Printf("cross product of $vtcXunit and $vtcYunit is %+v aka $vtcZunit\n", vtcXunit.CrossProduct(vtcYunit))
	fmt.Printf("cosinus of angle between $vtc1 and $vtc2 is %+v\n", vtc1.Cos(vtc2))
	fmt.Printf("cosinus of angle between $vtcXunit and $vtcYunit is %+v\n", vtcXunit.Cos(vtcYunit))
}
