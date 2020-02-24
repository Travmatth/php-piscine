package Vertex

import (
	"fmt"
	"io/ioutil"
	Color "d06/ex00"
)

var Verbose = false

type Vertex struct {
	x       float64
	y       float64
	z       float64
	w       float64
	color   *Color.Color
}

/* Constructor */

func NewVertex(arr []float64, color *Color.Color) (v *Vertex) {
	if len(arr) == 3 {
		arr = append(arr, 1.0)
	}	else if len(arr) != 4 {
		panic(fmt.Errorf("Incorrect Vertex init"))
	}
	if color == nil {
		color = Color.NewColor([]int{255, 255, 255})
	}
	v = &Vertex{x: arr[0], y: arr[1], z: arr[2], w: arr[3], color: color}
	if Verbose {
		fmt.Printf("%+v constructed\n", v)
	}
	return
}

/* Golang doesnt have destructors */

func (v *Vertex) Release() {
	if Verbose {
		fmt.Printf("%+v destructed\n", v)
	}
}
/* Getters */

func (v *Vertex) GetX() float64 {
	return v.x
}

func (v *Vertex) GetY() float64 {
	return v.y
}

func (v *Vertex) GetZ() float64 {
	return v.z
}

func (v *Vertex) GetW() float64 {
	return v.w
}

func (v *Vertex) GetColor() *Color.Color {
	return v.color
}

/* Setters */

func (v *Vertex) SetX(x float64) {
	v.x = x
}

func (v *Vertex) SetY(y float64) {
	v.y = y
}

func (v *Vertex) SetZ(z float64) {
	v.z = z
}

func (v *Vertex) SetW(w float64) {
	v.w = w
}

func (v *Vertex) SetColor(color *Color.Color) {
	v.color = color
}

/* Doc */

func  Doc() (doc string, err error) {
	b, err := ioutil.ReadFile("ex01/Vertex.doc.txt")
	if err != nil {
		return "", err
	}
	doc = string(b)
	return
}

func (v *Vertex) String() (repr string) {
	repr = fmt.Sprintf("Vertex( x:%.2f, y:%.2f, z:%.2f, w:%.2f ", v.x, v.y, v.z, v.w)
	if Verbose {
		repr += v.color.String() + " )"
	} else {
		repr += ")"
	}
	return
}