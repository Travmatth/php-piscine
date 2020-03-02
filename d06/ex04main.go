package main

import (
	Vertex "d06/ex01"
	Vector "d06/ex02"
	Matrix "d06/ex03"
	Camera "d06/ex04"
	"fmt"
	"math"
)

func main() {
	Vertex.Verbose = false
	Vector.Verbose = false
	Matrix.Verbose = false
	doc, err := Camera.Doc()
	if err == nil {
		fmt.Print(doc)
	}
	Camera.Verbose = true
	vtx0 := Vector.NewVector([]float{20.0, 20.0, 80.0})
	R := Matrix.NewMatrix(Matrix.RY, 0.0, math.Pi, nil, 0.0, 0.0, 0.0, 0.0)
	cam := Camera.NewCamera(vtx0, R, 640, 480, 60, 1.0, 100.0)
	ft.Println(cam)
	return
}
