package main

import (
	Ex04 "d07/ex04"
)

func main() {
	j := Ex04.Jamie{}
	c := Ex04.Cersei{}
	t := Ex04.Tyrion{}
	s := Ex04.Sansa{}

	j.SleepWith(t) //not even if im drunk
	j.SleepWith(s) // Lets do this
	j.SleepWith(c) // with pleasure, but only in a tower

	t.SleepWith(j) //not even if im drunk
	t.SleepWith(s) // Lets do this
	t.SleepWith(c) //not even if im drunk

}
