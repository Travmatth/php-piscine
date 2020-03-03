package Ex04

import "fmt"

type Jamie struct {
}

func (j *Jamie) SleepWith(p interface{}) {
	if _, ok := p.(Tyrion); ok == true {
		fmt.Println("Not even if I'm drunk !")
	} else if _, ok = p.(Sansa); ok == true {
		fmt.Println("Let's do this.")
	} else if _, ok = p.(Cersei); ok == true {
		fmt.Println("With pleasure, but only in a tower in Winterfell, then.")
	}
}
