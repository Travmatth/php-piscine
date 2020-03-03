package Camera

import (
	Vertex "d06/ex01"
	Matrix "d06/ex03"
	"io/ioutil"
)

type Camera struct {
	origin     *Vertex.Vertex
	tT         *Matrix.Matrix
	tR         *Matrix.Matrix
	projection *Matrix.Matrix
	width      float64
	height     float64
	ratio      float64
	fov        float64
	near       float64
	far        float64
}

func NewCamera(origin *Vertex.Vertex, tT *Matrix.Matrix, width, height, ratio, fov, near, far float64) (c *Camera) {
	tR := Matrix.NewMatrix(Matrix.TRANSLATION, 0, 0, nil, fov, ratio, near, far)
	projection := Matrix.NewMatrix(Matrix.PROJECTION, 0, 0, origin.Opposite(), 0, 0, 0, 0)
	c = &Camera{
		origin, tT, tR, projection, width, height, ratio, fov, near, far,
	}
	return
}

var Verbose = false

func (c *Camera) transpose(matrix *Matrix.Matrix) (m [16]float64) {
	orig := matrix.GetMatrix()
	for i := 0; i < 16; i++ {
		x, y := (i % 4), (i / 4)
		j := (x * 4) + y
		m[i], m[j] = orig[j], orig[i]
	}
	return
}

/* Transform world coordinates vertex into screen coordinates vertex */

func (c *Camera) WatchVertex(worldVertex *Vertex.Vertex) (v *Vertex.Vertex) {
	v = c.tR.Transform(c.tR.Transform(worldVertex))
	v.SetX(v.GetX() * c.ratio)
	v.SetY(v.GetY())
	v.SetColor(worldVertex.GetColor())
	return
}

/* Get Docs */

func Doc() (doc string, err error) {
	b, err := ioutil.ReadFile("ex06/Camera.doc.txt")
	if err != nil {
		return "", err
	}
	doc = string(b)
	return
}
