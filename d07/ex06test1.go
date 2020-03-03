package main

import (
	Ex06 "d07/ex06"
	"fmt"
)

type Footsoldier struct {
	name string
}

func MakeFootsoldier() *Footsoldier {
	return &Footsoldier{name: "foot soldier"}
}

func (f *Footsoldier) Fight(target string) {
	fmt.Printf("* draws his sword and runs towards %s *\n", target)
}

func (f *Footsoldier) GetName() (n string) {
	n = f.name
	return
}

type Archer struct {
	name string
}

func MakeArcher() *Footsoldier {
	return &Footsoldier{name: "archer"}
}

func (a *Archer) Fight(target string) {
	fmt.Printf("* shoots arrows at %s *\n", target)
}

func (a *Archer) GetName() (n string) {
	n = a.name
	return
}

type Assassin struct {
	name string
}

func MakeAssassin() *Footsoldier {
	return &Footsoldier{name: "assassin"}
}

func (a *Assassin) Fight(target string) {
	fmt.Printf("* creeps behind %s and stabs at it *\n", target)
}

func (a *Assassin) GetName() (n string) {
	n = a.name
	return
}

type Llama struct {
}

func MakeLlama() *Llama {
	return &Llama{}
}

func Fight(target string) {
	fmt.Printf("* spits at %s *\n", target)
}

func main() {
	uf := Ex06.MakeUnholyFactory()
	uf.Absorb(MakeFootsoldier())
	uf.Absorb(MakeFootsoldier())
	uf.Absorb(MakeArcher())
	uf.Absorb(MakeAssassin())
	uf.Absorb(MakeLlama())

	requestedFighters := []string{
		"foot soldier",
		"llama",
		"foot soldier",
		"archer",
		"foot soldier",
		"assassin",
		"foot soldier",
		"archer",
	}

	actualFighters := make([]Ex06.Fighter, 0)
	for _, rf := range requestedFighters {
		if f := uf.Fabricate(rf); f != nil {
			actualFighters = append(actualFighters, f)
		}
	}

	targets := []string{"the Hound", "Tyrion", "Podrick"}
	for _, f := range actualFighters {
		for _, t := range targets {
			f.Fight(t)
		}
	}
}
