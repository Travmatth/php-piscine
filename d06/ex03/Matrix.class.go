package Matrix

import (
	Vertex "d06/ex01"
	Vector "d06/ex02"
	"fmt"
	"io/ioutil"
	"math"
)

const (
	IDENTITY    = "IDENTITY"
	SCALE       = "SCALE"
	RX          = "0x ROTATION"
	RY          = "0y ROTATION"
	RZ          = "0z ROTATION"
	TRANSLATION = "TRANSLATION"
	PROJECTION  = "PROJECTION"
)

var Verbose = false

type Matrix struct {
	preset string
	scale  float64
	angle  float64
	vtc    *Vector.Vector
	fov    float64
	ratio  float64
	near   float64
	far    float64
	matrix [16]float64
}

func NewMatrix(preset string, scale, angle float64, vtc *Vector.Vector, fov, ratio, near, far float64) (m *Matrix) {
	var matrix [16]float64
	for i := 0; i < 16; i++ {
		matrix[i] = 0
	}
	m = &Matrix{preset, scale, angle, vtc, fov, ratio, near, far, matrix}
	m.identity(1)
	switch preset {
	case IDENTITY:
	case SCALE:
		m.identity(m.scale)
	case RX:
		m.rotateX()
	case RY:
		m.rotateY()
	case RZ:
		m.rotateZ()
	case TRANSLATION:
		m.translation()
	case PROJECTION:
		m.projection()
	}
	if Verbose && m.preset == IDENTITY {
		fmt.Printf("Matrix %s instance constructed\n", m.preset)
	} else if Verbose {
		fmt.Printf("Matrix %s preset instance constructed\n", m.preset)
	}
	return
}

func (m *Matrix) GetMatrix() [16]float64 {
	return m.matrix
}

func (m *Matrix) copy() (c *Matrix) {
	var matrix [16]float64
	for i := 0; i < 16; i++ {
		matrix[i] = 0
	}
	c = &Matrix{m.preset, m.scale, m.angle, m.vtc, m.fov, m.ratio, m.near, m.far, matrix}
	return
}

func (m *Matrix) Mult(rhs *Matrix) (c *Matrix) {
	var mult [16]float64
	c = m.copy()
	for y := 0; y < 16; y += 4 {
		for x := 0; x < 4; x += 1 {
			mult[y+x] = 0
			mult[y+x] += m.matrix[y] * rhs.matrix[x]
			mult[y+x] += m.matrix[y+1] * rhs.matrix[x+4]
			mult[y+x] += m.matrix[y+2] * rhs.matrix[x+8]
			mult[y+x] += m.matrix[y+3] * rhs.matrix[x+12]
		}
	}
	c.matrix = mult
	return
}

func (m *Matrix) Transform(v *Vertex.Vertex, x, y, z, w int) (t float64) {
	xPos := v.GetX() * m.matrix[x]
	yPos := v.GetY() * m.matrix[y]
	zPos := v.GetZ() * m.matrix[z]
	wPos := v.GetW() * m.matrix[w]
	t = xPos + yPos + zPos + wPos
	return
}

func (m *Matrix) TransformVertex(v *Vertex.Vertex) (next *Vertex.Vertex) {
	coords := make([]float64, 4)
	coords[0] = m.Transform(v, 0, 1, 2, 3)
	coords[1] = m.Transform(v, 4, 5, 6, 7)
	coords[2] = m.Transform(v, 8, 9, 10, 11)
	coords[3] = m.Transform(v, 12, 13, 14, 15)
	next = Vertex.NewVertex(coords, v.GetColor())
	return
}

func (m *Matrix) identity(scale float64) {
	m.matrix[0] = scale
	m.matrix[5] = scale
	m.matrix[10] = scale
	m.matrix[15] = 1
}

func (m *Matrix) translation() {
	m.matrix[3] = m.vtc.GetX()
	m.matrix[7] = m.vtc.GetY()
	m.matrix[11] = m.vtc.GetZ()
}

func (m *Matrix) projection() {
	m.matrix[5] = 1 / math.Tan(0.5*(m.fov*(math.Pi/180)))
	m.matrix[0] = m.matrix[5] / m.ratio
	m.matrix[10] = -(-m.near - m.far) / (m.near - m.far)
	m.matrix[14] = -1
	m.matrix[11] = (2 * m.near * m.far) / (m.near - m.far)
	m.matrix[15] = 0
}

func (m *Matrix) rotateX() {
	m.matrix[0] = 1
	m.matrix[5] = math.Cos(m.angle)
	m.matrix[6] = -math.Sin(m.angle)
	m.matrix[9] = math.Sin(m.angle)
	m.matrix[10] = math.Cos(m.angle)
}

func (m *Matrix) rotateY() {
	m.matrix[0] = math.Cos(m.angle)
	m.matrix[2] = math.Sin(m.angle)
	m.matrix[5] = 1
	m.matrix[8] = -math.Sin(m.angle)
	m.matrix[10] = math.Cos(m.angle)
}

func (m *Matrix) rotateZ() {
	m.identity(1)
	m.matrix[0] = math.Cos(m.angle)
	m.matrix[1] = -math.Sin(m.angle)
	m.matrix[4] = math.Sin(m.angle)
	m.matrix[5] = math.Cos(m.angle)
	m.matrix[10] = 1
}

func (m *Matrix) Release() {
	if Verbose {
		fmt.Println("Matrix instance destructed")
	}
}

func fmtRowString(c string, m [16]float64, start, end int) (s string) {
	args := []interface{}{c}
	for i := start; i < end; i++ {
		args = append(args, m[i])
	}
	s = fmt.Sprintf("%s | %.2f | %.2f | %.2f | %.2f\n", args...)
	return
}

func (m *Matrix) String() (s string) {
	s = "M | vtcX | vtcY | vtcZ | vtxO\n"
	s += "-----------------------------\n"
	s += fmtRowString("x", m.matrix, 0, 4)
	s += fmtRowString("y", m.matrix, 4, 8)
	s += fmtRowString("z", m.matrix, 8, 12)
	s += fmtRowString("w", m.matrix, 12, 16)
	return
}

/* Get Docs */

func Doc() (doc string, err error) {
	b, err := ioutil.ReadFile("ex03/Matrix.doc.txt")
	if err != nil {
		return "", err
	}
	doc = string(b)
	return
}
