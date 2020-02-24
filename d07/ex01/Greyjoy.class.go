package Greyjoy

type Greyjoy struct {
	familyMotto string
}

func NewGreyjoy() (g *Greyjoy) {
	g = &Greyjoy{"We do not sow"}
	return
}

func (g *Greyjoy) GetFamilyMotto() (s string) {
	s = g.familyMotto
	return
}
