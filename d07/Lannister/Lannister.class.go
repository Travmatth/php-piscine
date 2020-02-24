package Lannister

import "fmt"

type Lannister struct {
}

func NewLannister() (l *Lannister) {
	fmt.Println("A Lannister is born !")
	l = &Lannister{}
	return
}

func (l *Lannister) GetSize() (size string) {
	return "Average"
}

func (l *Lannister) HouseMotto() (m string) {
	m = "Hear me roar!"
	return
}
