package Color

import (
	"fmt"
	"io/ioutil"
)

var Verbose = false

type Color struct {
	red   uint8
	blue  uint8
	green uint8
}

/* Constructor */

func NewColor(arr []int) (c *Color) {
	c = nil
	if len(arr) == 1 {
		val := arr[0]
		c = &Color{
			red:   uint8(val),
			green: uint8(val >> 8),
			blue:  uint8(val >> 16),
		}
	} else if len(arr) == 3 {
		c = &Color{
			red:   uint8(arr[0]),
			green: uint8(arr[1]),
			blue:  uint8(arr[2]),
		}
	} else {
		panic(fmt.Errorf("Incorrect Color init"))
	}
	if Verbose {
		fmt.Printf("%+v constructed\n", c)
	}
	return
}

/* Golang doesn't have desctructors */

func (c *Color) Release() {
	if Verbose {
		fmt.Printf("%+v destructed\n", c)
	}
}
func (c *Color) String() string {
	return fmt.Sprintf("Color( red: %3d, green: %3d, blue: %3d )", c.red, c.green, c.blue)
}

func ColorVerbose(new bool) {
	Verbose = new
}

func Doc() (doc string, err error) {
	b, err := ioutil.ReadFile("ex00/Color.doc.txt")
	if err != nil {
		return "", err
	}
	doc = string(b)
	return
}

func (c *Color) Add(rhs *Color) (new *Color) {
	new = &Color{
		red:   c.red + rhs.red,
		green: c.green + rhs.green,
		blue:  c.blue + rhs.blue,
	}
	if Verbose {
		fmt.Printf("%+v constructed\n", new)
	}
	return
}

func (c *Color) Sub(rhs *Color) (new *Color) {
	new = &Color{
		red:   c.red - rhs.red,
		green: c.green - rhs.green,
		blue:  c.blue - rhs.blue,
	}
	if Verbose {
		fmt.Printf("%+v constructed\n", new)
	}
	return
}

func (c *Color) Mult(k float64) (new *Color) {
	new = &Color{
		red:   uint8(float64(c.red) * k),
		green: uint8(float64(c.green) * k),
		blue:  uint8(float64(c.blue) * k),
	}
	if Verbose {
		fmt.Printf("%+v constructed\n", new)
	}
	return
}
