package Vector

import (
	"math"
	"fmt"
	"io/ioutil"
	Vertex "d06/ex01"
)

var Verbose = false

type Vector struct {
	x float64
	y float64
	z float64
	w float64
}

/* Constructors */

func NewVector(orig *Vertex.Vertex, dest *Vertex.Vertex) (v *Vector) {
	if orig == nil {
		orig = Vertex.NewVertex([]float64{0, 0, 0, 1,}, nil)
	} else if dest == nil {
		panic("Incorrect Vector Init")
	}
	v = &Vector{
		x: dest.GetX() - orig.GetX(),
		y: dest.GetY() - orig.GetY(),
		z: dest.GetZ() - orig.GetZ(),
		w: 0,
	}
	if Verbose {
		fmt.Printf("%+v constructed\n", v)
	}
	return
}

/* Golang doesnt have Destructors */

func (v *Vector) Release() {
	fmt.Printf("%+v destructed\n", v)
}

/* Vector operations */

func (v *Vector) Magnitude() (m float64) {
	m = (v.x * v.x) + (v.y * v.y) + (v.z * v.z)
	m = float64(math.Sqrt(float64(m)))
	return
}

func (v *Vector) Normalize() (n *Vector) {
	len := v.Magnitude()
	if len == 1 {
		coords := []float64{v.x, v.y, v.z}
		n = NewVector(nil, Vertex.NewVertex(coords, nil))
	} else {
		coords := []float64{v.x / len, v.y / len, v.z / len}
		n = NewVector(nil, Vertex.NewVertex(coords, nil))
	}
	return
}

func (v *Vector) Add(rhs *Vector) (a *Vector) {
	coords := []float64{v.x + rhs.x, v.y + rhs.y, v.z + rhs.z}
	a = NewVector(nil, Vertex.NewVertex(coords, nil))
	return
}

func (v *Vector) Sub(rhs *Vector) (s *Vector) {
	coords := []float64{v.x - rhs.x, v.y - rhs.y, v.z - rhs.z}
	s = NewVector(nil, Vertex.NewVertex(coords, nil))
	return
}

func (v *Vector) Opposite() (o *Vector) {
	coords := []float64{-v.x, -v.y, -v.z}
	o = NewVector(nil, Vertex.NewVertex(coords, nil))
	return
}

func (v *Vector) ScalarProduct(k float64) (s *Vector) {
	coords := []float64{v.x * k, v.y * k, v.z * k}
	s = NewVector(nil, Vertex.NewVertex(coords, nil))
	return
}

func (v *Vector) DotProduct(rhs *Vector) (d float64) {
	d = (v.x * rhs.x) + (v.y * rhs.y) + (v.z * rhs.z)
	return
}

func (v *Vector) Cos(rhs *Vector) (c float64) {
	numerator := (v.x * rhs.x) + (v.y * rhs.y) + (v.z * rhs.z)
	f := func(x float64) float64 {return x * x}
	denominator := (f(v.x) + f(v.y) + f(v.z)) * (f(rhs.x) + f(rhs.y) + f(rhs.z))
	c = numerator / math.Sqrt(denominator);
	return
}

func (v *Vector) CrossProduct(rhs *Vector) (c *Vector) {
	coords := []float64{
		(v.y * rhs.GetZ()) - (v.z * rhs.GetY()),
		(v.z * rhs.GetX()) - (v.x * rhs.GetZ()),
		(v.x * rhs.GetY()) - (v.y * rhs.GetX()),
	}
	c = NewVector(nil, Vertex.NewVertex(coords, nil))
	return
}

/* String representation */

func (v *Vector) String() (s string) {
	s = fmt.Sprintf("Vector( x:%.2f, y:%.2f, z:%.2f, w:%.2f )", v.x, v.y, v.z, v.w)
	return
}

/* Get Docs */

func  Doc() (doc string, err error) {
	b, err := ioutil.ReadFile("ex02/Vector.doc.txt")
	if err != nil {
		return "", err
	}
	doc = string(b)
	return
}

/* Setters */

func (v *Vector) SetX(x float64) {
	v.x = x
}

func (v *Vector) SetY(y float64) {
	v.y = y
}

func (v *Vector) SetZ(z float64) {
	v.z = z
}

func (v *Vector) SetW(w float64) {
	v.w = w
}

/* Getters */

func (v *Vector) GetX() float64 {
	return v.x
}

func (v *Vector) GetY() float64 {
	return v.y
}

func (v *Vector) GetZ() float64 {
	return v.z
}

func (v *Vector) GetW() float64 {
	return v.w
}
