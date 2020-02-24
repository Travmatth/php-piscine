package main

import (
	Vertex "d06/ex01"
	Vector "d06/ex02"
	Matrix "d06/ex03"
	"fmt"
)

func main() {
	Vertex.Verbose = false
	Vector.Verbose = false

	doc, err := Matrix.Doc()
	if err == nil {
		fmt.Print(doc)
	}
	Matrix.Verbose = true

	fmt.Println("Let's start with an harmless identity matrix :")
	I := Matrix.NewMatrix(Matrix.IDENTITY, 0, 0, nil, 0, 0, 0, 0)
	defer I.Release()

	fmt.Printf("%+v\n", I)

	fmt.Println("So far, so good. Let's create a translation matrix now.")
	vtx := Vertex.NewVertex([]float64{20.0, 20.0, 0.0}, nil)
	vtc := Vector.NewVector(nil, vtx)
	T := Matrix.NewMatrix(Matrix.TRANSLATION, 0, 0, vtc, 0, 0, 0, 0)
	defer T.Release()
	fmt.Println(T)

	fmt.Println("A scale matrix is no big deal.")
	S := Matrix.NewMatrix(Matrix.SCALE, 10.0, 0, vtc, 0, 0, 0, 0)
	defer S.Release()
	fmt.Println(S)

	fmt.Println("A Rotation along the OX axis :")
	RX := Matrix.NewMatrix(Matrix.RX, 0, 0.78539816339744830962, nil, 0, 0, 0, 0)
	defer RX.Release()
	fmt.Println(RX)

	fmt.Println("Or along the OY axis :")
	RY := Matrix.NewMatrix(Matrix.RY, 0, 1.57079632679489661923, nil, 0, 0, 0, 0)
	defer RY.Release()
	fmt.Println(RY)

	fmt.Println("Do a barrel roll !")
	RZ := Matrix.NewMatrix(Matrix.RZ, 0, 2*3.14159265358979323846, nil, 0, 0, 0, 0)
	defer RZ.Release()
	fmt.Println(RZ)

	fmt.Println("The bad guy now, the projection matrix : 3D to 2D !")
	fmt.Println("The values are arbitray. We'll decipher them in the next exercice.")
	P := Matrix.NewMatrix(Matrix.PROJECTION, 0, 0, nil, 60, 640.0/480.0, 1.0, -50.0)
	defer P.Release()
	fmt.Println(P)

	fmt.Println("Matrices are so awesome, that they can be combined !")
	fmt.Println("This is a model matrix that scales, then rotates around OY axis,")
	fmt.Println("then rotates around OX axis and finally translates.")
	fmt.Println("Please note the reverse operations order. It's not an error.")
	M := T.Mult(RX).Mult(RY).Mult(S)
	defer M.Release()
	fmt.Println(M)

	fmt.Println("What can you do with a matrix and a vertex ?")
	vtxA := Vertex.NewVertex([]float64{1.0, 1.0, 0.0}, nil)
	fmt.Println(vtxA)
	fmt.Println(M)
	fmt.Println("Transform the damn vertex !")
	vtxB := M.TransformVertex(vtxA)
	fmt.Println(vtxB)
}
