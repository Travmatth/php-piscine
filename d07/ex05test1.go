package main

import (
	Ex05 "d07/ex05"
	"fmt"
)

type JonSnow struct {
}

func (j *JonSnow) Fight() {
	fmt.Println("* looses his wolf on the enemy, and charges with courage *")
}

func (j *JonSnow) IsABastard() bool {
	return true
}

type MaesterAemon struct {
}

func (m *MaesterAemon) SendRavens() {
	fmt.Println("* sends a raven carrying an important message *")
}

type SamwellTarly struct {
}

func (s *SamwellTarly) Fight() {
	fmt.Println("* flees, finds a girl, grows a spine, and defends her to the bitter end *")
}

func main() {
	jon := &JonSnow{}
	aemon := &MaesterAemon{}
	sam := &SamwellTarly{}
	nw := &Ex05.NightsWatch{}

	nw.Recruit(jon)
	nw.Recruit(aemon)
	nw.Recruit(sam)
	nw.Fight()
}
