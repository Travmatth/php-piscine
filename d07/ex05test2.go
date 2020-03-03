package main

import (
	Ex05 "d07/ex05"
	"fmt"
)

type Varys struct {
}

func (c *Varys) PretendToFight() {
	fmt.Println("* finds someone to fight for him *")
}

func main() {
	varys := &Varys{}
	nw := &Ex05.NightsWatch{}
	nw.Recruit(varys)
	nw.Fight()
}
