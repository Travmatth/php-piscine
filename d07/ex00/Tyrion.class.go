package Tyrion

import (
	"d07/Lannister"
)

type Tyrion struct {
	Lannister.Lannister
}

func NewTyrion() (t *Tyrion) {
	t = &Tyrion{}
	return
}

func (t *Tyrion) GetSize() (size string) {
	return "Short"
}
