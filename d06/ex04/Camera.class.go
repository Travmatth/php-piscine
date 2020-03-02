package Camera

import (
	Vertex "d06/ex01"
	Matrix "d06/ex03"
	"io/ioutil"
)

type Camera struct {
	origin      *Vertex.Vertex
	orientation *Matrix.Matrix
	projection  *Matrix.Matrix
	width       float64
	height      float64
	ratio       float64
	fov         float64
	near        float64
	far         float64
}

func NewCamera(origin *Vertex.Vertex, orientation *Matrix.Matrix, width, height, ratio, fov, near, far float64) (c *Camera) {
	projection := Matrix.NewMatrix(Matrix.PROJECTION, 0, 0, nil, fov, ratio, near, far)
	c = &Camera{
		origin, orientation, projection, width, height, ratio, fov, near, far,
	}
	return
}

var Verbose = false

func (c *Camera) transpose() (m [16]float64) {
	var matrix [4][4]float64
	original := c.projection.GetMatrix()
	for i := 0; i < 16; i++ {
		x, y := i%4, i/4
		matrix[y][x] = original[i]
	}
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			matrix[x][y], matrix[y][x] = matrix[y][x], matrix[x][y]
		}
	}
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			m[(y*4)+x] = matrix[y][x]
		}
	}
	return
}

/* Transform world coordinates vertex into screen coordinates vertex */

func (c *Camera) WatchVertex(worldVertex *Vertex.Vertex) (v *Vertex.Vertex) {
	v = c.proj.transform(c.tr.transform(worldVertex))
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
