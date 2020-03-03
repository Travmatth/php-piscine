package Ex06

import "fmt"

type UnholyFactory struct {
	fighters map[string]Fighter
}

func MakeUnholyFactory() (u *UnholyFactory) {
	fighters := make(map[string]Fighter)
	u = &UnholyFactory{fighters}
	return
}

func (u *UnholyFactory) Absorb(soldier interface{}) {
	var ok bool
	var fighter Fighter
	if fighter, ok = soldier.(Fighter); !ok {
		fmt.Println("(Factory can't absob this, it's not a fighter)")
	} else if _, ok = u.fighters[fighter.GetName()]; ok {
		fmt.Printf("*Factory has already absorbed a fighter of type %s)\n", fighter.GetName())
	} else {
		u.fighters[fighter.GetName()] = fighter
		fmt.Printf("*Factory absorbed fighter of type %s)\n", fighter.GetName())
	}
}

func (u *UnholyFactory) Fabricate(fighter string) (f Fighter) {
	var ok bool
	f, ok = u.fighters[fighter]
	if ok {
		fmt.Printf("(Factory fabricates a fighter of type %s)\n", fighter)
	} else {
		fmt.Printf("(Factory hasn't absorbed any fighter of type %s)\n", fighter)
	}
	return
}
