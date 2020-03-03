package Ex05

type NightsWatch struct {
	fighters []IFighter
}

func (n *NightsWatch) Recruit(i interface{}) {
	if fighter, ok := i.(IFighter); ok == true {
		n.fighters = append(n.fighters, fighter)
	}
}

func (n *NightsWatch) Fight() {
	for _, fighter := range n.fighters {
		fighter.Fight()
	}
}
