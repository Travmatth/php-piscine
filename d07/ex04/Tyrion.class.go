package Ex04

import "fmt"

type Tyrion struct {
}

func (t *Tyrion) SleepWith(p interface{}) {
	if _, ok := p.(Jamie); ok == true {
		fmt.Println("Not even if I'm drunk !")
	} else if _, ok = p.(Sansa); ok == true {
		fmt.Println("Let's do this.")
	} else if _, ok = p.(Cersei); ok == true {
		fmt.Println("Not even if I'm drunk !")
	}
}
