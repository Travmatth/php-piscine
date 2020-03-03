package main

import (
	Ex06 "d07/ex06"
)

type CrippledSoldier struct {
	name string
}

func MakeCrippledSoldier() (c *CrippledSoldier) {
	c = &CrippledSoldier{name: "crippled soldier"}
	return
}

func main() {
	uf := Ex06.MakeUnholyFactory()
	uf.Absorb(MakeCrippledSoldier())

	requestedFighters := []string{
		"crippled soldier",
		"crippled soldier",
		"crippled soldier",
		"crippled soldier",
	}

	actualFighters := []Ex06.Fighter{}
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
